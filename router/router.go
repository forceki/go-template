package router

import (
	"github.com/forceki/invent-be/src/auth"
	"github.com/forceki/invent-be/src/checkin"
	"github.com/forceki/invent-be/src/gudang"
	"github.com/forceki/invent-be/src/items"
	"github.com/forceki/invent-be/src/kategori"
	"github.com/forceki/invent-be/src/supplier"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, app *fiber.App) {
	api := app.Group("/api")

	items.ItemsRouter(db, api)
	gudang.GudangRouter(db, api)
	supplier.SupplierRouter(db, api)
	checkin.CheckinRouter(db, api)
	kategori.KategoriRouter(db, api)
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
