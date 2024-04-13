package middlewareHandlers

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(success bool, message string, data interface{}) Response {
	return Response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func SendResponse(c *fiber.Ctx, response Response, status int) error {
	return c.Status(status).JSON(response)
}
