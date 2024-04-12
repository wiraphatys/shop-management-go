package adminHandlers

import "github.com/gofiber/fiber/v2"

type AdminHandler interface {
	GetAdminByEmail(c *fiber.Ctx) error
	CreateAdmin(c *fiber.Ctx) error
}
