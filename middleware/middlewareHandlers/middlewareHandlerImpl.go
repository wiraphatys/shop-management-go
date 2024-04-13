package middlewareHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/shop-management-go/middleware/middlewareUsecases"
)

type middlewareHandlerImpl struct {
	middlewareUsecase middlewareUsecases.MiddlewareUsecase
}

func NewMiddlewareHandler(middlewareUsecase middlewareUsecases.MiddlewareUsecase) MiddlewareHandler {
	return &middlewareHandlerImpl{
		middlewareUsecase: middlewareUsecase,
	}
}

func (h *middlewareHandlerImpl) Authenticated(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	if accessToken == "" {
		response := NewResponse(false, "Unauthorized", nil)
		return SendResponse(c, response, fiber.StatusUnauthorized)
	}

	user, err := h.middlewareUsecase.VerifyToken(accessToken)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response, fiber.ErrBadRequest.Code)
	}

	c.Locals("user", user)
	return c.Next()
}
