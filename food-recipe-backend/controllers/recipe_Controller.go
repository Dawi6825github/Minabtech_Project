package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"backend/internal/recipes"
	"backend/models"
)

// CreateRecipe adds a new recipe
func CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := recipes.CreateRecipe(&recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create recipe"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Recipe added successfully"})
}

func GetAllRecipes(c *gin.Context) {
	recipes, err := recipes.GetAllRecipes() // Assuming this function exists
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

// GetRecipe retrieves a recipe by ID
func GetRecipe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	recipe, err := recipes.GetRecipeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

	// Register the new route for getting all recipes
	r.GET("/api/recipes", GetAllRecipes) // Add this line to your router setup
func UpdateRecipe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var recipe models.Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := recipes.UpdateRecipe(id, &recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update recipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe updated successfully"})
}

// DeleteRecipe removes a recipe
func DeleteRecipe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := recipes.DeleteRecipe(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete recipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted successfully"})
}
