package orderHandlers

import "github.com/gofiber/fiber/v2"

type OrderHandler interface {
	GetAllOrders(c *fiber.Ctx) error
	GetOrderById(c *fiber.Ctx) error
	CreateOrder(c *fiber.Ctx) error
	UpdateOrderLineById(c *fiber.Ctx) error
	DeleteOrderById(c *fiber.Ctx) error
	DeleteOrderLineById(c *fiber.Ctx) error
}
