package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// User structure
type User struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// JWT secret key
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// SignupHandler registers a new user
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Store user in DB (dummy implementation)
	user.ID = "user_" + fmt.Sprintf("%d", time.Now().Unix())

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// LoginHandler authenticates a user and returns JWT token
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Fetch user from DB (dummy data)
	dummyUser := User{
		ID:       "user_12345",
		Email:    "test@example.com",
		Password: "$2a$14$9W9OJ1eXkgO/zQ1u.TFvvOpwILYXfM9CWOBk9Mj5uFqxgy/JgYomK", // hashed password
	}

	// Validate email
	if user.Email != dummyUser.Email {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Compare hashed passwords
	err := bcrypt.CompareHashAndPassword([]byte(dummyUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": dummyUser.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// ProfileHandler returns user profile
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	claims := jwt.MapClaims{}

	// Parse JWT token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Fetch user info (dummy data)
	userProfile := User{
		ID:       "user_12345",
		FullName: "John Doe",
		Email:    "test@example.com",
	}

	// Return user profile
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userProfile)
}
