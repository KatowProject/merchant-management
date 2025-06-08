package user

import (
	"merchant-management/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Response struct {
	Status 	string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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
		c.JSON(http.StatusInternalServerError, ErrorResponse{
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
	id := c.Param("id")

	user := User{ID: id, Name: "Dummy", Email: "dummy@example.com"}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.ID = "3"
	c.JSON(http.StatusCreated, newUser)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser.ID = id
	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "User deleted", "id": id})
}
