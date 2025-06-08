package user

import (
	"merchant-management/internal/db"
	"merchant-management/internal/repository"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {

    db, db_err := db.GetDB()
    if db_err != nil {
        panic("Failed to connect to the database: " + db_err.Error())
    }

    userRepo := repository.NewUserRepository(db)
    handler := NewUserHandler(userRepo)

    userGroup := rg.Group("/users")
    {
        userGroup.GET("/", handler.GetAllUsers)
        userGroup.GET("/:id", handler.GetUserByID)
        userGroup.POST("/", handler.CreateUser)
        userGroup.PUT("/:id", handler.UpdateUser)
        userGroup.DELETE("/:id", handler.DeleteUser)
    }
}
