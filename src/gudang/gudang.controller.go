package gudang

import (
	"github.com/forceki/invent-be/handler"
	"github.com/gofiber/fiber/v2"
)

type gudangController struct {
	gudangService GudangService
}

func NewGudangController(gudangService GudangService) *gudangController {
	return &gudangController{gudangService: gudangService}
}

func (c *gudangController) FindAll(f *fiber.Ctx) error {
	data, err := c.gudangService.FindAll()

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)

	}

	return handler.ResponseHttp(f, 200, 1, "success", data)

}

func (c *gudangController) Master(f *fiber.Ctx) error {

	data, err := c.gudangService.FindAll()

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	type res struct {
		Id    int    `json:"id"`
		Label string `json:"label"`
	}

	var supp []res
	for _, item := range *data {
		if item.Status == true {
			key := res{
				Id:    item.GudangId,
				Label: item.Nama,
			}
			supp = append(supp, key)
		}

	}

	return handler.ResponseHttp(f, 200, 1, "success", supp)
}

func (c *gudangController) Create(f *fiber.Ctx) error {
	var body gudangRequest
	err := f.BodyParser(&body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	errors := handler.ValidateStruct(&body)

	if errors != nil {
		return handler.ResponseHttp(f, 501, 0, "error", errors)
	}

	err = c.gudangService.Create(body)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "created", nil)
}

func (c *gudangController) Update(f *fiber.Ctx) error {
	var body gudangRequest
	Id := f.Query("id")
	err := f.BodyParser(&body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}
	errors := handler.ValidateStruct(&body)

	if errors != nil {
		return handler.ResponseHttp(f, 501, 0, "error", errors)
	}
	err = c.gudangService.Update(Id, body)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "updated", nil)

}

func (c *gudangController) Delete(f *fiber.Ctx) error {
	Id := f.Query("id")

	err := c.gudangService.Delete(Id)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "deleted", nil)

}
