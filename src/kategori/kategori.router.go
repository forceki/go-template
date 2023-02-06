package kategori

import (
	"github.com/forceki/invent-be/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func KategoriRouter(db *gorm.DB, router fiber.Router) {
	repository := NewKategoriRepository(db)
	service := NewKategoriService(repository)
	controller := NewKategoriController(service)

	app := router.Group("/kategori", middleware.Auth)

	app.Get("/", controller.FindAll)
	app.Post("/", controller.Create)
	app.Put("/", controller.Update)
	app.Delete("/", controller.Delete)
}
