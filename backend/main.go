package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	client := connectToMongoDB()
	defer disconnectFromMongoDB(client)
	startServer()
}

func connectToMongoDB() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to MongoDB!")
	return client
}

func disconnectFromMongoDB(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}

func startServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(os.Getenv("PORT"), r)
}
