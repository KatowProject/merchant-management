package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		log.Println("Received ping request")

		c.JSON(200, gin.H{
			"message": "pong",
			"status":  200,
		})
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	r.Run(":" + PORT)
}
