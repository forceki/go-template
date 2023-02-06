package handler

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseHttp(f *fiber.Ctx, httpCode int, status int, message string, Data interface{}) error {
	response := Response{
		Status:  status,
		Message: message,
		Data:    Data,
	}

	return f.Status(httpCode).JSON(response)
}
