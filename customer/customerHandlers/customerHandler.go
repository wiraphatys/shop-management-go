package customerHandlers

import "github.com/gofiber/fiber/v2"

type CustomerHandler interface {
	GetAllCustomers(c *fiber.Ctx) error
	GetCustomerByEmail(c *fiber.Ctx) error
	RegisterCustomer(c *fiber.Ctx) error
	UpdateCustomerByEmail(c *fiber.Ctx) error
	DeleteCustomerByEmail(c *fiber.Ctx) error
}
