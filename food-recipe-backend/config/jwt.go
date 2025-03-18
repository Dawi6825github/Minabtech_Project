package config

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims struct for JWT
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(config.Config.JWTExpirationHours) * time.Hour)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.JWTSecret))
	if err != nil {
		log.Printf("Error generating JWT token: %v", err)
		return "", err
	}

	return tokenString, nil
}

// VerifyToken validates the JWT token and returns claims
func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
