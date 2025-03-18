package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection instance
var DB *gorm.DB

// ConnectDB initializes the PostgreSQL database connection
func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		Config.DBHost, Config.DBUser, Config.DBPassword, Config.DBName, Config.DBPort, Config.DBSSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("✅ Connected to the PostgreSQL database successfully")
}
