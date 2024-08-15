package storage

import (
	"EasyProcess/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// --- User ---

func (db Storage) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	singleResult := db.UserCollection.FindOne(*db.Context, bson.D{{Key: "email", Value: email}})
	err := singleResult.Decode(&user)
	// Check if got results otherwise user wont be nil
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &user, err
}

func (db Storage) GetUserForApiKey(apiKey string) (*models.User, error) {
	var user models.User
	singleResult := db.UserCollection.FindOne(*db.Context, bson.D{{Key: "apiKey", Value: apiKey}})
	err := singleResult.Decode(&user)
	// Check if got results otherwise user wont be nil
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &user, err
}
