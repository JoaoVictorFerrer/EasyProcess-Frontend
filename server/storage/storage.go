package storage

import (
	"EasyProcess/models"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	Client         *mongo.Client
	Context        *context.Context
	UserCollection *mongo.Collection
}

type StorageInterface interface {
	// --- READ ---

	// User
	GetUserByEmail(email string) (*models.User, error)
	GetUserForApiKey(apiKey string) (*models.User, error)

	// --- CREATE ---

	// User
	CreateUser(user *models.User) error

	// --- UPDATE ---

	// ..

	// --- DELETE ---

	// User
	DeleteUserByEmail(email string) error
}

func InitStorage(mongoClient *mongo.Client) Storage {
	var databaseName string
	if os.Getenv("ENV") == "DEV" {
		databaseName = "development"
	} else {
		databaseName = "production"
	}
	context := context.Background()
	return Storage{
		Client:         mongoClient,
		Context:        &context,
		UserCollection: mongoClient.Database(databaseName).Collection("user"),
	}
}
