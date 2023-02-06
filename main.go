package main

import (
	"github.com/forceki/invent-be/database"
	"github.com/forceki/invent-be/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	db := database.ConnectDB()

	// Listen on PORT 3000
	app.Use(cors.New())
	router.SetupRoutes(db, app)

	app.Listen(":3000")

}
