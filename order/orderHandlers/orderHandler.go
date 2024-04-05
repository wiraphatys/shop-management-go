package orderHandlers

import "github.com/gofiber/fiber/v2"

type OrderHandler interface {
	CreateOrder(c *fiber.Ctx) error
}
