package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type File struct {
	ID    string
	Nonce []byte
	Salt  []byte
}

func ConnectToMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to PlanetScale!")
	return db, nil
}

func UploadToDB(db *sql.DB, id string, nonce []byte, salt []byte) error {
	query := `INSERT INTO files (id, nonce, salt) VALUES (?, ?, ?)`
	_, err := db.Exec(query, id, nonce, salt)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
		return err
	}
	log.Println("Data successfully inserted!")
	return nil
}

func CheckIfExists(db *sql.DB, id string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM files WHERE id=?)`
	var exists bool
	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if data exists: %v", err)
		return false, err
	}
	return exists, nil
}

func GetFileInfo(db *sql.DB, id string, password string) (*File, error) {
	query := `SELECT id, nonce, salt FROM files WHERE id=?`
	row := db.QueryRow(query, id)

	var file File
	err := row.Scan(&file.ID, &file.Nonce, &file.Salt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			log.Fatalf("Failed to get data: %v", err)
			return nil, err
		}
	}
	return &file, nil
}
