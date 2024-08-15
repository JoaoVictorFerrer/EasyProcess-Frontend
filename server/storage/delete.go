package storage

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (db Storage) DeleteUserByEmail(email string) error {
	// Delete user account
	_, err := db.UserCollection.DeleteOne(*db.Context, bson.D{{Key: "email", Value: email}})
	return err
}
