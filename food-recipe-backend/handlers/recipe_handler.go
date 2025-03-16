package handlers

import (
	"backend/models"
	"github.com/username/food-recipe-backend/models"	"net/http"
	"github.com/gin-gonic/gin"

)

var recipes []models.Recipe

func GetRecipes(c *gin.Context) {
    c.JSON(http.StatusOK, recipes)
}

func CreateRecipe(c *gin.Context) {
    var newRecipe models.Recipe
    if err := c.ShouldBindJSON(&newRecipe); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    recipes = append(recipes, newRecipe)
    c.JSON(http.StatusCreated, newRecipe)
}
