package routes

import (
    "food-recipe-backend/handlers"
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    
    // Recipe Routes
    api.GET("/recipes", handlers.GetRecipes)
    api.POST("/recipes", handlers.CreateRecipe)
    
    // User Routes
    api.POST("/register", handlers.RegisterUser)
    api.POST("/login", handlers.LoginUser)
}
