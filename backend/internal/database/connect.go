package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

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
