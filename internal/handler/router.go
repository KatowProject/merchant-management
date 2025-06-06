package handler

import (
	"github.com/gin-gonic/gin"
	"merchant-management/internal/handler/user"
)

func NewRouter() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api")
    {
        user.RegisterUserRoutes(api)
    }

    return r
}