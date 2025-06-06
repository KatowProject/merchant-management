package user

import (
    "github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
    handler := NewUserHandler()

    userGroup := rg.Group("/users")
    {
        userGroup.GET("/", handler.GetAllUsers)
        userGroup.GET("/:id", handler.GetUserByID)
        userGroup.POST("/", handler.CreateUser)
        userGroup.PUT("/:id", handler.UpdateUser)
        userGroup.DELETE("/:id", handler.DeleteUser)
    }
}
