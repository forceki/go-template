package checkin

import (
	"strconv"

	"github.com/forceki/invent-be/handler"
	"github.com/gofiber/fiber/v2"
)

type checkinController struct {
	checkinService CheckinService
}

func NewCheckinController(checkinService CheckinService) *checkinController {
	return &checkinController{checkinService}
}

func (c *checkinController) Create(f *fiber.Ctx) error {
	var body CheckinRes
	err := f.BodyParser(&body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)

	}

	errors := handler.ValidateStruct(&body)
	if errors != nil {
		return handler.ResponseHttp(f, 501, 0, "error", errors)
	}

	err = c.checkinService.Create(body)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "created", nil)
}

func (c *checkinController) FindAll(f *fiber.Ctx) error {
	Status := f.Query("status")
	intVar, _ := strconv.Atoi(Status)
	data, err := c.checkinService.FindAll(intVar)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 200, 1, "success", data)

}

func (c *checkinController) Delete(f *fiber.Ctx) error {
	Id := f.Query("id")

	err := c.checkinService.Delete(Id)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 200, 1, "deleted", nil)
}

func (c *checkinController) FindOne(f *fiber.Ctx) error {
	Id := f.Query("id")

	data, err := c.checkinService.FindOne(Id)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}
	return handler.ResponseHttp(f, 200, 1, "success", data)
}

func (c *checkinController) Update(f *fiber.Ctx) error {
	var body CheckinRes
	Id := f.Query("id")
	err := f.BodyParser(&body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)

	}

	errors := handler.ValidateStruct(&body)
	if errors != nil {
		return handler.ResponseHttp(f, 501, 0, "error", errors)
	}

	err = c.checkinService.Update(Id, body)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "updated", nil)
}
