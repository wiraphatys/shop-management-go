package productHandlers

import "github.com/gofiber/fiber/v2"

type ProductHandler interface {
	GetAllProducts(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProductById(c *fiber.Ctx) error
	DeleteProductById(c *fiber.Ctx) error
}
