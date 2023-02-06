package gudang

import (
	"github.com/forceki/invent-be/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GudangRouter(db *gorm.DB, router fiber.Router) {
	repository := NewGudangRepository(db)
	service := NewGudangService(repository)
	controller := NewGudangController(service)

	app := router.Group("/gudang", middleware.Auth)

	app.Get("/", controller.FindAll)
	app.Get("/master", controller.Master)
	app.Post("/", controller.Create)
	app.Put("/", controller.Update)
	app.Delete("/", controller.Delete)
}
