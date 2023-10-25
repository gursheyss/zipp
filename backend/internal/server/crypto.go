package server

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"zipp/internal/database"

	"github.com/xdg-go/pbkdf2"
)

func encryptFile(file multipart.File, password string) ([]byte, []byte, []byte, error) {
	plaintext, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, nil, err
	}

	salt := make([]byte, 8)
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

	// Convert hex strings back to byte slices
	nonce, err := hex.DecodeString(string(fileInfo.Nonce))
	if err != nil {
		return nil, fmt.Errorf("failed to decode nonce: %w", err)
	}

	salt, err := hex.DecodeString(string(fileInfo.Salt))
	if err != nil {
		return nil, fmt.Errorf("failed to decode salt: %w", err)
	}

	key := pbkdf2.Key([]byte(password), salt, 4096, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Strip the nonce from the ciphertext
	ciphertext := unencryptedFile[len(nonce):]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed, possibly due to incorrect password: %w", err)
	}

	return plaintext, nil
}
