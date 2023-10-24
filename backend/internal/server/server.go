package server

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"zipp/internal/database"
	"zipp/internal/s3"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/xdg-go/pbkdf2"
)

func StartServer(db *sql.DB) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handleRoot)
	r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		handleUpload(w, r, db)
	})
	r.Get("/download", func(w http.ResponseWriter, r *http.Request) {
		handleDownload(w, r, db)
	})
	r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
		handleCheck(w, r, db)
	})
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatalf("Server failed %v", err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hey lol"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleCheck(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := r.URL.Query().Get("id")
	exists, err := database.CheckIfExists(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(strconv.FormatBool(exists)))
}

func handleDownload(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := r.URL.Query().Get("id")
	password := r.URL.Query().Get("password")
	ctx := context.Background()

	client, err := s3.ConnectToAWS(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	unencryptedFiles, err := s3.DownloadAllFiles(ctx, client, "zipp-files", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	decryptedFiles := make(map[string][]byte)
	for key, unencryptedFile := range unencryptedFiles {
		decryptedFile, err := decryptFile(unencryptedFile, db, id, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		decryptedFiles[key] = decryptedFile
	}

	jsonData, err := json.Marshal(decryptedFiles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func handleUpload(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	id := r.FormValue("id")
	password := r.FormValue("password")

	log.Println("Unique ID: ", id)
	log.Println("Received password: ", password)
	log.Println("Filename: ", header.Filename)
	log.Println("Size: ", header.Size)
	log.Println("Content Type: ", header.Header.Get("Content-Type"))

	ctx := context.Background()
	client, err := s3.ConnectToAWS(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	salt, nonce, ciphertext, err := encryptFile(file, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = database.UploadToDB(db, id, nonce, salt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ciphertextReader := bytes.NewReader(ciphertext)

	key, err := s3.UploadFile(ctx, client, ciphertextReader, "zipp-files", header.Filename, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	keyJson, err := json.Marshal(map[string]string{"key": key})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(keyJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func encryptFile(file multipart.File, password string) ([]byte, []byte, []byte, error) {
	plaintext, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, nil, err
	}

	salt := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, nil, nil, err
	}

	salt = make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, nil, nil, err
	}
	key := pbkdf2.Key([]byte(password), salt, 4096, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return salt, nonce, ciphertext, nil
}

func decryptFile(unencryptedFile []byte, db *sql.DB, id string, password string) ([]byte, error) {
	fileInfo, err := database.GetFileInfo(db, id, password)
	if err != nil {
		return nil, err
	}

	key := pbkdf2.Key([]byte(password), fileInfo.Salt, 4096, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := gcm.Open(nil, fileInfo.Nonce, unencryptedFile, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed, possibly due to incorrect password: %w", err)
	}

	return plaintext, nil
}
