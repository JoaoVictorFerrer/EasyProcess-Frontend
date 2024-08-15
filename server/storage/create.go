package storage

import (
	"EasyProcess/models"
)

// --- User ---

func (db Storage) CreateUser(user *models.User) error {
	_, err := db.UserCollection.InsertOne(*db.Context, user)
	return err
}
