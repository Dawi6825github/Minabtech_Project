package handlers

import (
	"encoding/json"
	"../models""net/http"
	"github.com/dawit6825github/food-recipe-backend/services"	

	"github.com/gorilla/mux"
)

// Add a recipe to the user's bookmark list
func BookmarkRecipe(w http.ResponseWriter, r *http.Request) {
	var bookmark models.Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if bookmark.UserID == 0 || bookmark.RecipeID == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Save bookmark to the database
	bookmark, err := services.AddBookmark(bookmark)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookmark)
}
// Get all bookmarks for a user
func GetUserBookmarks(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]

	// Fetch all bookmarks for the user
	bookmarks, err := services.GetBookmarksByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the list of bookmarks
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookmarks)
}

// Remove a recipe from the user's bookmark list
func RemoveBookmark(w http.ResponseWriter, r *http.Request) {
	bookmarkID := mux.Vars(r)["bookmarkID"]

	// Delete bookmark
	err := services.DeleteBookmark(bookmarkID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusNoContent)
}
