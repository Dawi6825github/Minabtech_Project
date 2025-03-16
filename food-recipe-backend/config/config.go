package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
        os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"))

    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to the database!")
}
