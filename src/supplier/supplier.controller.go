package supplier

import (
	"github.com/forceki/invent-be/handler"
	"github.com/gofiber/fiber/v2"
)

type supplierController struct {
	supplierService SupplierService
}

func NewSupplierController(supplierService SupplierService) *supplierController {
	return &supplierController{supplierService: supplierService}
}

func (c *supplierController) FindAll(f *fiber.Ctx) error {

	data, err := c.supplierService.FindAll()

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 200, 1, "success", data)
}

func (c *supplierController) Master(f *fiber.Ctx) error {

	data, err := c.supplierService.FindAll()

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	type res struct {
		Id    int    `json:"id"`
		Label string `json:"label"`
	}

	var supp []res
	for _, item := range *data {
		key := res{
			Id:    item.ID,
			Label: item.Nama,
		}

		supp = append(supp, key)
	}

	return handler.ResponseHttp(f, 200, 1, "success", supp)
}

func (c *supplierController) Create(f *fiber.Ctx) error {
	var body SupplierRequest
	err := f.BodyParser(&body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)

	}

	errors := handler.ValidateStruct(&body)
	if errors != nil {
		return handler.ResponseHttp(f, 501, 0, "error", errors)
	}

	err = c.supplierService.Create(body)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "created", nil)
}

func (c *supplierController) Update(f *fiber.Ctx) error {
	var body SupplierRequest
	Id := f.Query("id")
	err := f.BodyParser(&body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	errors := handler.ValidateStruct(&body)

	if errors != nil {
		return handler.ResponseHttp(f, 501, 0, "error", errors)
	}

	err = c.supplierService.Update(Id, body)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "updated", nil)
}

func (c *supplierController) Delete(f *fiber.Ctx) error {
	Id := f.Query("id")

	err := c.supplierService.Delete(Id)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "deleted", nil)
}
