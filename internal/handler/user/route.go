package user

import (
	"merchant-management/internal/db"
	"merchant-management/internal/repository"
	"merchant-management/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {

    db, err := db.GetDB()
    if err != nil {
        panic("Failed to connect to the database: " + err.Error())
    }

    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    handler := NewUserHandler(userService)

    userGroup := rg.Group("/users")
    {
        userGroup.GET("/", handler.GetAllUsers)
        userGroup.GET("/:id", handler.GetUserByID)
        userGroup.POST("/", handler.CreateUser)
        userGroup.PUT("/:id", handler.UpdateUser)
        userGroup.DELETE("/:id", handler.DeleteUser)
    }
}
