package utils

import (
	"context"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/iamNanak/CineGenie/Server/StreamServer/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SignedDetails struct {
	Email     string
	UserID    string
	Role      string
	FirstName string
	LastName  string
	jwt.RegisteredClaims
}

var SecretKey string = os.Getenv("JWT_SECRET_KEY")
var userCollection *mongo.Collection = database.OpenCollection("users")

func GenerateAccessAndRefreshTokens(email, firstName, lastName, role, userId string) (string, string, error) {
	// Access Token
	accessClaims := &SignedDetails{
		Email:     email,
		UserID:    userId,
		Role:      role,
		FirstName: firstName,
		LastName:  lastName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "CineGenie",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Refresh Token
	refreshClaims := &SignedDetails{
		Email:     email,
		UserID:    userId,
		Role:      role,
		FirstName: firstName,
		LastName:  lastName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "CineGenie",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	signedAccessToken, err := accessToken.SignedString([]byte(SecretKey))
	if err != nil {
		return "", "", err
	}

	signedRefreshToken, err := refreshToken.SignedString([]byte(SecretKey))
	if err != nil {
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil
}

func UpdateTokens(signedAccessToken, signedRefreshToken, userId string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

}
