package auth

import (
	"merchant-management/internal/db"
	"merchant-management/internal/repository"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	db, err := db.GetDB()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	useRepo := repository.NewUserRepository(db)
	authHandler := NewAuthHandler(useRepo)

	auth := router.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}
}
