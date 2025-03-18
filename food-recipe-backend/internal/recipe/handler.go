package recipes

import (
	"net/http"

	"food-recipe-backend/internal/models"
	"food-recipe-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RecipeHandler defines the structure for the handler
type RecipeHandler struct {
	Service *RecipeService
}

// NewRecipeHandler creates a new handler instance for recipes
func NewRecipeHandler(service *RecipeService) *RecipeHandler {
	return &RecipeHandler{Service: service}
}

// GetAllRecipes handles the GET request to fetch all recipes
func (h *RecipeHandler) GetAllRecipes(c *gin.Context) {
	recipes, err := h.Service.GetAll()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to fetch recipes")
		return
	}
	utils.RespondJSON(c, http.StatusOK, recipes)
}

// CreateRecipe handles the POST request to create a new recipe
func (h *RecipeHandler) CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdRecipe, err := h.Service.Create(&recipe)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create recipe")
		return
	}

	utils.RespondJSON(c, http.StatusCreated, createdRecipe)
}

// UpdateRecipe handles the PUT request to update an existing recipe
func (h *RecipeHandler) UpdateRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	updatedRecipe, err := h.Service.Update(recipe.ID, &recipe)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update recipe")
		return
	}

	utils.RespondJSON(c, http.StatusOK, updatedRecipe)
}

// DeleteRecipe handles the DELETE request to delete a recipe by ID
func (h *RecipeHandler) DeleteRecipe(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.Delete(id); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete recipe")
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Recipe deleted successfully"})
}
