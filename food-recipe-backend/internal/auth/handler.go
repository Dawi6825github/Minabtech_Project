package auth

import (
	"encoding/json"
	"net/http"
	"food-recipe-backend/internal/utils"
)

// LoginHandler handles user login and returns a JWT token
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials UserCredentials

	// Decode the request body to get login credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Authenticate the user and generate a token
	token, err := AuthenticateUser(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Respond with the generated JWT token
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

// SignupHandler handles user signup and returns a JWT token
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var credentials UserCredentials

	// Decode the request body to get signup credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Signup the user and generate a token
	token, err := SignupUser(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with the generated JWT token
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}
