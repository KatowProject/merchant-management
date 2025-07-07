package user

import (
	"merchant-management/internal/model"
	"merchant-management/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)
type Response struct {
	Status 	string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type UserHandler struct{
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUser()
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
	var uri struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Invalid user ID",
		})
		return
	}

	print("id:", id.ID)

	user, err := h.service.GetUserByID(id)
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

	createdUser, err := h.service.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "Failed to create user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "User created successfully",
		Data:    createdUser,
	})

}

func (h *UserHandler) UpdateUser(c *gin.Context) {

}

func (h *UserHandler) DeleteUser(c *gin.Context) {

}
