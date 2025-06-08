package handler

import (
	"merchant-management/internal/handler/user"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api")
    {   
        user.RegisterUserRoutes(api)
    }

    return r
}