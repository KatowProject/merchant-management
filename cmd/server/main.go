package main

import (
	"log"
	"os"

	"merchant-management/internal/db"
	"merchant-management/internal/handler"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	_, db_err := db.Connect()
	if db_err != nil {
		log.Fatalf("Failed to connect to the database: %v", db_err)
	}

	router := handler.NewRouter()

	PORT := os.Getenv("APP_PORT")
	if PORT == "" {
		PORT = "8080"
	}

	router.Run(":" + PORT)
}
