package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Password     *string            `json:"password" validate:"required,min=6"`
	Username     *string            `json:"username" validate:"required"`
	Token        *string            `json:"token"`
	User_type    *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Created_at   time.Time          `json:"created_at"`
	Updated_at   time.Time          `json:"updated_at"`
	Token_expiry time.Time          `json:"token_expiry"`
	User_id      string             `json:"user_id"`
}
