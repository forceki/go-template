package auth

import (
	"github.com/forceki/invent-be/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRouter(db *gorm.DB, router fiber.Router) {
	repository := NewAuthRepository(db)
	service := NewAuthService(repository)
	controller := NewAuthController(service)

	app := router.Group("/auth")

	app.Post("/register", middleware.Auth, controller.Create)
	app.Post("/login", controller.Login)
}
