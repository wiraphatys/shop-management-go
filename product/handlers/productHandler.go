package handlers

import "github.com/gofiber/fiber/v2"

type ProductHandler interface {
	GetAllProducts(c *fiber.Ctx) error
}
