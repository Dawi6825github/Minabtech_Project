package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte("your_secret_key_here") // This should be stored securely

// JWTClaims represents the structure of the JWT payload
type JWTClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

// ParseJWT parses and validates the JWT token
func ParseJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

// GenerateJWT generates a new JWT token for a user
func GenerateJWT(userID uint64) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
			Issuer:    "your_project_name",
		},
	}

	// Creating the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
