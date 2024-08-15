package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Surname   string             `bson:"surname" json:"surname"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Role      string             `bson:"role" json:"role"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updateAt" json:"updatedAt"`
}

// Add created at or updated at dates to model when being modified
func (user *User) MarshalBSON() ([]byte, error) {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	user.UpdatedAt = time.Now()
	type my User
	return bson.Marshal((*my)(user))
}
