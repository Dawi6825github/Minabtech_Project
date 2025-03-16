package main

import (
	"backend/golang/handlers"
	"backend/golang/middleware"
	"backend/golang/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	db := utils.ConnectDB()
	defer db.Close()

	// Setup Router
	router := mux.NewRouter()

	// Middleware
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.CORSMiddleware)

	// Routes
	router.HandleFunc("/signup", handlers.Signup).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/recipes", handlers.GetRecipes).Methods("GET")

	// Start Server
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
