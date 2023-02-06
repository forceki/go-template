package auth

import (
	"github.com/forceki/invent-be/handler"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authService AuthService
}

func NewAuthController(authService AuthService) *authController {
	return &authController{authService}
}

func (c *authController) Create(f *fiber.Ctx) error {
	var body UsersRes

	err := f.BodyParser(&body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)

	}

	errors := handler.ValidateStruct(&body)
	if errors != nil {
		return handler.ResponseHttp(f, 501, 0, "error", errors)
	}

	err = c.authService.Create(body)
	if err != nil {
		return handler.ResponseHttp(f, 501, 0, err.Error(), nil)
	}

	return handler.ResponseHttp(f, 201, 1, "created", nil)
}

func (c *authController) Login(f *fiber.Ctx) error {
	type login struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var body login

	err := f.BodyParser(&body)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, "username or password cannot match", nil)
	}

	errors := handler.ValidateStruct(&body)
	if errors != nil {
		return handler.ResponseHttp(f, 501, 0, "error", errors)
	}

	data, err := c.authService.Login(body.Username, body.Password)

	if err != nil {
		return handler.ResponseHttp(f, 501, 0, "username or password cannot match", nil)
	}

	return handler.ResponseHttp(f, 201, 1, "Logged", data)
}
