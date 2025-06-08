package db

import (
	"fmt"
	"log"
	"merchant-management/internal/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB

func Connect() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}

	models := []interface{}{
		&model.User{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatal("Failed to auto-migrate database:", err)
		return nil, err
	}

	instance = db

	log.Println("Database connection established successfully")

	return db, nil
}

func GetDB() (*gorm.DB, error) {
	if instance == nil {
		return nil, fmt.Errorf("database connection is not established")
	}
	return instance, nil
}
