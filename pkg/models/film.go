// models/film.go
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Film struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title" json:"title" validate:"required"`
	Director    string             `bson:"director" json:"director" validate:"required"`
	ReleaseDate time.Time          `bson:"releaseDate" json:"releaseDate" validate:"required"`
	Cast        []string           `bson:"cast" json:"cast" validate:"required"`
	Genre       []string           `bson:"genre" json:"genre" validate:"required"`
	Synopsis    string             `bson:"synopsis" json:"synopsis" validate:"required"`
	User_ID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	Created_by  string             `bson:"created_by" json:"created_by"`
}
