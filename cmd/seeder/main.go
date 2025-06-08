package main

import (
	"merchant-management/internal/db"
	"merchant-management/seeder"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	database, err := db.Connect()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	seeder.RunAll(database)
}

