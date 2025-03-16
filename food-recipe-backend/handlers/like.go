package handlers

import (
	"net/http"
	"food-recipe-app/golang/models"
	"food-recipe-app/golang/services"
	"github.com/gorilla/mux"
	"encoding/json"
)

// Like a recipe
func LikeRecipe(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if like.RecipeID == "" || like.UserID == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Add the like to the database
	like, err := services.AddLike(like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the like data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(like)
}

// Unlike a recipe
func UnlikeRecipe(w http.ResponseWriter, r *http.Request) {
	likeID := mux.Vars(r)["likeID"]

	// Remove the like from the database
	err := services.RemoveLike(likeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusNoContent)
}

// Get all likes for a recipe
func GetRecipeLikes(w http.ResponseWriter, r *http.Request) {
	recipeID := mux.Vars(r)["recipeID"]

	// Fetch likes for the recipe
	likes, err := services.GetLikesByRecipeID(recipeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the list of likes
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(likes)
}
