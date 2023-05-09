package router

import (
	"github.com/forceki/invent-be/src/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, app *fiber.App) {
	api := app.Group("/api")

	auth.AuthRouter(db, api)

	api.Get("/test", testin)
}

func testin(f *fiber.Ctx) error {
	return f.Status(201).JSON(fiber.Map{
		"status":  1,
		"message": "found",
		"data":    nil,
	})
}
