package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	id        primitive.ObjectID `bson:"_id,omitempty"`
	discordId string             `bson:"discordId,omitempty"`
	username  string             `bson:"username,omitempty"`
	createdAt time.Time          `bson:"createdAt,omitempty"`
	updatedAt time.Time          `bson:"updatedAt,omitempty"`
}

type UserRepository interface{}
