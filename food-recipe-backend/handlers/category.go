package handlers

import (
    "encoding/json"
    "net/http"
    "your_project/golang/models"
    "your_project/golang/services"
    "your_project/golang/utils"
    "github.com/gorilla/mux"
)

// CreateCategory handles the creation of a new category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
    var category models.Category
    // Parse the request body into category struct
    err := json.NewDecoder(r.Body).Decode(&category)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid input")
        return
    }

    // Call the service to create the category in the DB
    err = services.CreateCategory(&category)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create category")
        return
    }

    // Respond with success
    utils.RespondWithJSON(w, http.StatusCreated, category)
}

// GetCategories retrieves all categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
    categories, err := services.GetAllCategories()
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch categories")
        return
    }

    // Respond with the list of categories
    utils.RespondWithJSON(w, http.StatusOK, categories)
}

// GetCategoryByID retrieves a category by its ID
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    categoryID := vars["id"]

    category, err := services.GetCategoryByID(categoryID)
    if err != nil {
        utils.RespondWithError(w, http.StatusNotFound, "Category not found")
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, category)
}

// UpdateCategory updates a category by its ID
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    categoryID := vars["id"]

    var category models.Category
    // Parse the request body into category struct
    err := json.NewDecoder(r.Body).Decode(&category)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid input")
        return
    }

    // Call the service to update the category in the DB
    err = services.UpdateCategory(categoryID, &category)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update category")
        return
    }

    // Respond with success
    utils.RespondWithJSON(w, http.StatusOK, category)
}

// DeleteCategory deletes a category by its ID
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    categoryID := vars["id"]

    err := services.DeleteCategory(categoryID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete category")
        return
    }

    // Respond with success
    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Category deleted successfully"})
}
