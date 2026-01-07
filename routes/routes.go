package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-tutorial/controllers"
	"go-tutorial/utils"
	"time"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(app *fiber.App) {
	fmt.Println("Routes")

	// Home route (optional caching if you want)
	app.Get("/", controllers.GetHome)
	app.Get("/users",
		utils.CacheMiddleware(10*time.Minute, utils.StaticKey("users_html")),
		controllers.GetUsers,
	)
	app.Get("/api/users",
		utils.CacheMiddleware(5*time.Minute, utils.StaticKey("users_json")),
		controllers.GetUsersJSON,
	)
	// app.Post("/users", controllers.AddUser)
}
