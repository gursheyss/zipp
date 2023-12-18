package main

import (
	"log"
	"os"
	"strings"

	"zipp/internal/database"
	"zipp/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading environment variables from OS")
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			os.Setenv(pair[0], pair[1])
		}
	}

	db, err := database.ConnectToMySQL()
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	server.StartServer(db)
}
