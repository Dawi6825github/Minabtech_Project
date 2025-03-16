package handlers

import (
	"encoding/json"
	"food-recipe-app/golang/models"
	"food-recipe-app/golang/services"
	"net/http"

	"github.com/gorilla/mux"
)

// Rate a recipe
func RateRecipe(w http.ResponseWriter, r *http.Request) {
	var rating models.Rating
	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if rating.RecipeID == "" || rating.UserID == "" || rating.RatingValue <= 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Save rating to the database
	rating, err := services.AddRating(rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the rating details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rating)
}

// Get the average rating of a recipe
func GetRecipeRating(w http.ResponseWriter, r *http.Request) {
	recipeID := mux.Vars(r)["recipeID"]

	// Fetch the average rating
	averageRating, err := services.GetAverageRating(recipeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the average rating
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(averageRating)
}
