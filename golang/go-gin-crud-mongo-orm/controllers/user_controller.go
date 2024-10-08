// Package controllers contains the handler functions for various routes.
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gitish/polyglot_training/models"
	"github.com/gitish/polyglot_training/services"

	"github.com/gin-gonic/gin"
)

// UserController defines the interface for user-related API operations.
type UserController interface {
	CreateUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

// userControllerImpl implements the UserController interface.
type userControllerImpl struct {
	userService services.UserService
}

// NewUserController creates a new instance of userControllerImpl.
func NewUserController(userService services.UserService) UserController {
	return &userControllerImpl{userService: userService}
}

// CreateUser handles the POST request to create a user.
func (ctrl *userControllerImpl) CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	createdUser := ctrl.userService.CreateUser(newUser)
	c.JSON(http.StatusCreated, createdUser)
}

// GetUsers handles the GET request to retrieve all users.
func (ctrl *userControllerImpl) GetUsers(c *gin.Context) {
	users := ctrl.userService.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// GetUserByID handles the GET request to retrieve a user by ID.
func (ctrl *userControllerImpl) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, exists := ctrl.userService.GetUserByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser handles the PUT request to update a user by ID.
func (ctrl *userControllerImpl) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	updatedUser.ID = id
	user, updated := ctrl.userService.UpdateUser(updatedUser)
	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser handles the DELETE request to remove a user by ID.
func (ctrl *userControllerImpl) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	deleted := ctrl.userService.DeleteUser(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
