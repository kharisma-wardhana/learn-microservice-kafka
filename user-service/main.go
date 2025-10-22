package main

import (
	"context"
	"log"
	"os"

	db "github.com/kharisma-wardhana/learn-microservice-kafka/shared/db"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	postgresDB, err := db.NewPostgresDB(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer postgresDB.Close(context.Background())

	log.Println("Successfully connected to the database")
}
