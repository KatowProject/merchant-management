package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		log.Println("Received ping request")

		c.JSON(200, gin.H{
			"message": "pong",
			"status":  200,
		})
	})

	r.Run(":8080") // listen and serve on
}
