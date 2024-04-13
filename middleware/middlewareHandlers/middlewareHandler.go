package middlewareHandlers

import "github.com/gofiber/fiber/v2"

type MiddlewareHandler interface {
	Authenticated(c *fiber.Ctx) error
}
