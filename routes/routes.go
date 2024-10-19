package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-tutorial/controllers"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(app *fiber.App) {
	// User routes
	fmt.Println("Routes")
	app.Get("/", controllers.GetHome)        // Display users
	app.Get("/users", controllers.GetUsers)        // Display users
	app.Get("/api/users", controllers.GetUsersJSON)        // Display users
	// app.Post("/users", controllers.AddUser)   // Add a new user
}
