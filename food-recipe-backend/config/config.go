package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig holds application configuration values
type AppConfig struct {
	Port                string
	DBHost              string
	DBPort              string
	DBUser              string
	DBPassword          string
	DBName              string
	DBSSLMode           string
	JWTSecret           string
	JWTExpirationHours  int
	HasuraGraphQLURL    string
	HasuraAdminSecret   string
	ChapaSecretKey      string
}

// Config is a global variable holding all configuration values
var Config AppConfig

// LoadConfig loads environment variables from the .env file
func LoadConfig() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	Config = AppConfig{
		Port:               getEnv("PORT", "8080"),
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPassword:         getEnv("DB_PASSWORD", ""),
		DBName:             getEnv("DB_NAME", "app_db"),
		DBSSLMode:          getEnv("DB_SSLMODE", "disable"),
		JWTSecret:          getEnv("JWT_SECRET", "supersecret"),
		HasuraGraphQLURL:   getEnv("HASURA_GRAPHQL_URL", ""),
		HasuraAdminSecret:  getEnv("HASURA_ADMIN_SECRET", ""),
		ChapaSecretKey:     getEnv("CHAPA_SECRET_KEY", ""),
	}

	// Convert JWT expiration hours to an integer
	expirationHours, err := strconv.Atoi(getEnv("JWT_EXPIRATION_HOURS", "24"))
	if err != nil {
		log.Fatalf("Invalid JWT expiration hours: %v", err)
	}
	Config.JWTExpirationHours = expirationHours
}

// getEnv reads an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
