package auth

import (
	"errors"
	"fmt"
	"time"
	// "food-recipe-backend/internal/users"
	"food-recipe-backend/internal/utils"
	"github.com/dgrijalva/jwt-go"
)

// UserCredentials defines the expected structure for login/signup
type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// JWT claims structure (moved to utils package)

// JWT Secret Key (Ideally loaded from environment variables)
var secretKey = []byte("your_secret_key_here")

// GenerateJWT generates a new JWT token
func GenerateJWT(userID uint64) (string, error) {
	claims := utils.JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // 72-hour expiration
			Issuer:    "your_project_name",
		},
	}

	// Creating the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// AuthenticateUser validates user credentials and returns JWT token
func AuthenticateUser(credentials UserCredentials) (string, error) {
	user, err := users.GetUserByEmail(credentials.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !ValidatePassword(user.Password, credentials.Password) {
		return "", errors.New("incorrect password")
	}

	// Generate JWT token
	token, err := GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// SignupUser creates a new user and returns a JWT token
func SignupUser(credentials UserCredentials) (string, error) {
	// Hash the user's password
	hashedPassword, err := HashPassword(credentials.Password)
	if err != nil {
		return "", err
	}

	// Create a new user record in the database
	userID, err := users.CreateUser(credentials.Email, hashedPassword)
	if err != nil {
		return "", err
	}

	// Generate JWT token
	token, err := GenerateJWT(userID)
	if err != nil {
		return "", err
	}

	return token, nil
}
