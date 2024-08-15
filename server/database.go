package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase() *mongo.Client {
	dbConnectionURL := os.Getenv("MONGO_URL")
	if dbConnectionURL == "" {
		panic("Failed to extract dbConnectionURL from environment")
	}
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbConnectionURL))
	if err != nil {
		panic(err)
	}
	log.Printf("🗄️ Connected to database 🗄️")
	return mongoClient
}
