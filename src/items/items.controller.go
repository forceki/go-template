package items

import (
	handler "github.com/forceki/invent-be/handler"
	"github.com/gofiber/fiber/v2"
)

type itemsController struct {
	itemsService ItemsService
}

func NewItemController(itemsService ItemsService) *itemsController {
	return &itemsController{itemsService: itemsService}
}

func (c *itemsController) FindAll(f *fiber.Ctx) error {

	Id := f.Query("id")
	data, err := c.itemsService.FindAll(Id)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 200, 1, "success", data)

}

func (c *itemsController) Master(f *fiber.Ctx) error {

	Id := f.Query("id")
	data, err := c.itemsService.FindAll(Id)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	type res struct {
		Id   string `json:"id"`
		Nama string `json:"nama"`
	}

	var item []res
	for _, e := range *data {
		key := res{
			Id:   e.ItemId,
			Nama: e.Nama,
		}

		item = append(item, key)
	}

	return handler.ResponseHttp(f, 200, 1, "success", item)

}

func (c *itemsController) Create(f *fiber.Ctx) error {
	var body ItemsRequest
	err := f.BodyParser(&body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	errors := handler.ValidateStruct(&body)

	if errors != nil {
		return handler.ResponseHttp(f, 501, 0, "error", errors)
	}

	err = c.itemsService.Create(body)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "created", nil)

}

func (c *itemsController) Update(f *fiber.Ctx) error {
	var body ItemsRequest
	Id := f.Query("id")
	err := f.BodyParser(&body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	errors := handler.ValidateStruct(&body)

	if errors != nil {
		return handler.ResponseHttp(f, 501, 0, "error", errors)
	}

	err = c.itemsService.Update(Id, body)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "updated", nil)

}

func (c *itemsController) Delete(f *fiber.Ctx) error {
	Id := f.Query("id")

	err := c.itemsService.Delete(Id)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "deleted", nil)

}
