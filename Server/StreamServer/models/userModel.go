package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID    string        `bson:"user_id" json:"user_id" validate:"required"`
	FirstName string        `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	LastName  string        `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Password  string        `bson:"password" json:"password" validate:"required,min=6"`
	Email     string        `bson:"email" json:"email" validate:"required,email"`
	Role      string        `bson:"role" json:"role" validate:"required"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
	JWTToken  string        `bson:"jwt_token" json:"jwt_token"`
}
