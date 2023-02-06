package supplier

import (
	"github.com/forceki/invent-be/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SupplierRouter(db *gorm.DB, router fiber.Router) {
	repository := NewSupplierRepository(db)
	service := NewSupplierService(repository)
	controller := NewSupplierController(service)

	app := router.Group("/supplier", middleware.Auth)

	app.Get("/", controller.FindAll)
	app.Get("/master", controller.Master)
	app.Post("/", controller.Create)
	app.Put("/", controller.Update)
	app.Delete("/", controller.Delete)
}
