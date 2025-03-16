package handlers

import (
	"encoding/json"
	"food-recipe-app/golang/models"
	"food-recipe-app/golang/services"
	"net/http"

	"github.com/gorilla/mux"
)

// Create a comment on a recipe
func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if comment.RecipeID == "" || comment.UserID == "" || comment.Content == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Save the comment to the database
	comment, err := services.CreateComment(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the newly created comment
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
}

// Get all comments for a recipe
func GetComments(w http.ResponseWriter, r *http.Request) {
	recipeID := mux.Vars(r)["recipeID"]

	// Fetch comments for the recipe
	comments, err := services.GetCommentsByRecipeID(recipeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the comments
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

// Delete a comment by ID
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentID := mux.Vars(r)["commentID"]

	// Delete the comment
	err := services.DeleteComment(commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusNoContent)
}
