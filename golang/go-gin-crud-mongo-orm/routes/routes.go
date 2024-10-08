// Package routes defines the available routes for the application.
package routes

import (
	"github.com/gitish/polyglot_training/controllers"
	"github.com/gitish/polyglot_training/middleware"
	"github.com/gitish/polyglot_training/repositories"
	"github.com/gitish/polyglot_training/services"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the routes for the application.
func SetupRouter(client *mongo.Client) *gin.Engine {
	r := gin.Default()

	// Add the DurationMiddleware
	r.Use(middleware.DurationMiddleware())

	// Initialize repository, service, and controller
	userRepository := repositories.NewUserRepository(client)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	// User routes
	r.POST("/users", userController.CreateUser)
	r.GET("/users", userController.GetUsers)
	r.GET("/users/:id", userController.GetUserByID)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	return r
}
