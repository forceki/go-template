package checkin

import (
	"github.com/forceki/invent-be/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CheckinRouter(db *gorm.DB, router fiber.Router) {
	repository := NewCheckinRepository(db)
	service := NewCheckinService(repository)
	controller := NewCheckinController(service)

	app := router.Group("/checkin", middleware.Auth)

	app.Post("/", controller.Create)
	app.Get("/", controller.FindAll)
	app.Delete("/", controller.Delete)
	app.Get("/one", controller.FindOne)
	app.Put("/", controller.Update)
}
