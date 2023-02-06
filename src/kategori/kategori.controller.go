package kategori

import (
	"github.com/forceki/invent-be/handler"
	"github.com/gofiber/fiber/v2"
)

type kategoriController struct {
	kategoriService KategoriService
}

func NewKategoriController(kategoriService KategoriService) *kategoriController {
	return &kategoriController{kategoriService}
}

func (c *kategoriController) FindAll(f *fiber.Ctx) error {
	data, err := c.kategoriService.FindAll()

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 200, 1, "success", data)
}

func (c *kategoriController) Create(f *fiber.Ctx) error {
	var body Kategori
	err := f.BodyParser(&body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	err = c.kategoriService.Create(body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "created", nil)
}

func (c *kategoriController) Update(f *fiber.Ctx) error {
	var body Kategori
	Id := f.Query("id")
	err := f.BodyParser(&body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	err = c.kategoriService.Update(Id, body)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "updated", nil)
}

func (c *kategoriController) Delete(f *fiber.Ctx) error {
	Id := f.Query("id")

	err := c.kategoriService.Delete(Id)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "deleted", nil)

}
