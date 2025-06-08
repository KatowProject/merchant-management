package user

import (
	"merchant-management/internal/model"
	"merchant-management/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)
type Response struct {
	Status 	string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type UserHandler struct{
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "Failed to fetch users: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "Users fetched successfully",
		Data:    users,
	})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	var id uint
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Invalid user ID",
		})
		return
	}

	user, err := h.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "Failed to fetch user: " + err.Error(),
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "error",
			Message: "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "User fetched successfully",
		Data:    user,
	})

}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Invalid user data: " + err.Error(),
		})
		return
	}

}

func (h *UserHandler) UpdateUser(c *gin.Context) {

}

func (h *UserHandler) DeleteUser(c *gin.Context) {

}
