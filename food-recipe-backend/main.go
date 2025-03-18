package main

import (
	"log"
	"net/http"	
	"food-recipe-backend/config"
	"os"
    "food-recipe-backend/internal/utils"

	"food-recipe-backend/internal/auth"
	"food-recipe-backend/internal/recipe"
	"food-recipe-backend/internal/users"
	"food-recipe-backend/internal/middleware"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
	config.LoadEnv() // Load environment variables

	router := mux.NewRouter()

	// Middlewares
	router.Use(middleware.CORS)
	router.Use(middleware.Logging)

	// Authentication routes
	router.HandleFunc("/api/login", auth.LoginHandler).Methods("POST")
	router.HandleFunc("/api/register", auth.RegisterHandler).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(auth.JWTMiddleware) // Secure these routes
	protected.HandleFunc("/recipes", recipes.GetRecipes).Methods("GET")
	protected.HandleFunc("/users", users.GetUsers).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is running on port", port)
	config.ConnectDB() // Establish database connection
	log.Fatal(http.ListenAndServe(":"+port, router)) 
}
