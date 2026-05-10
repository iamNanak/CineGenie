package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID              bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID          string        `bson:"user_id" json:"user_id"`
	FirstName       string        `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	LastName        string        `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Password        string        `bson:"password" json:"password" validate:"required,min=6"`
	Email           string        `bson:"email" json:"email" validate:"required,email"`
	Role            string        `bson:"role" json:"role" validate:"oneof=ADMIN USER"`
	CreatedAt       time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time     `bson:"updated_at" json:"updated_at"`
	AccessToken     string        `bson:"access_token" json:"access_token"`
	RefreshToken    string        `bson:"refresh_token" json:"refresh_token"`
	FavouriteGenres []Genre       `bson:"favourite_genres" json:"favourite_genres" validate:"required,dive"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	UserID         string  `json:"user_id"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Email          string  `json:"email"`
	Role           string  `json:"role"`
	FavouriteGenre []Genre `json:"favourite_genres"`
	AccessToken    string  `json:"access_token"`
	RefreshToken   string  `json:"refresh_token"`
}
