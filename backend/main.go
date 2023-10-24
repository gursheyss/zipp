package main

import (
	"context"
	"log"

	"zipp/internal/database"
	"zipp/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	ctx := context.Background()
	client := database.ConnectToMongoDB(ctx)
	defer database.DisconnectFromMongoDB(ctx, client)
	server.StartServer()
}
