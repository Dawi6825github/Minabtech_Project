package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/models"
	"backend/utils"
)

// CreateRecipeHandler - Add a new recipe
func CreateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = models.CreateRecipe(&recipe)
	if err != nil {
		http.Error(w, "Error creating recipe", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, "Recipe created successfully", http.StatusCreated)
}

// GetRecipesHandler - Fetch all recipes
func GetRecipesHandler(w http.ResponseWriter, r *http.Request) {
	recipes, err := models.GetAllRecipes()
	if err != nil {
		http.Error(w, "Error fetching recipes", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, recipes, http.StatusOK)
}

// GetRecipeByIDHandler - Fetch a single recipe
func GetRecipeByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	recipe, err := models.GetRecipeByID(id)
	if err != nil {
		http.Error(w, "Recipe not found", http.StatusNotFound)
		return
	}

	utils.JSONResponse(w, recipe, http.StatusOK)
}

// UpdateRecipeHandler - Update a recipe
func UpdateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = models.UpdateRecipe(&recipe)
	if err != nil {
		http.Error(w, "Error updating recipe", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, "Recipe updated successfully", http.StatusOK)
}

// DeleteRecipeHandler - Delete a recipe
func DeleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	err := models.DeleteRecipe(id)
	if err != nil {
		http.Error(w, "Error deleting recipe", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, "Recipe deleted successfully", http.StatusOK)
}
