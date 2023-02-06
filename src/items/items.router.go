package items

import (
	middleware "github.com/forceki/invent-be/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ItemsRouter(db *gorm.DB, router fiber.Router) {
	repository := NewItemsRepository(db)
	service := NewItemsService(repository)
	controller := NewItemController(service)

	app := router.Group("/items", middleware.Auth)

	app.Get("/", controller.FindAll)
	app.Get("/master", controller.Master)
	app.Post("/", controller.Create)
	app.Put("/", controller.Update)
	app.Delete("/", controller.Delete)
}
