package authHandlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/shop-management-go/admin/adminEntities"
	"github.com/wiraphatys/shop-management-go/auth/authUsecases"
	"github.com/wiraphatys/shop-management-go/config"
)

type authHandlerImpl struct {
	authUsecase authUsecases.AuthUsecase
}

func NewAuthHandler(authUsecase authUsecases.AuthUsecase) AuthHandler {
	return &authHandlerImpl{
		authUsecase: authUsecase,
	}
}

func (h *authHandlerImpl) SignIn(c *fiber.Ctx) error {
	reqBody := new(adminEntities.AdminData)
	if err := c.BodyParser(reqBody); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	token, err := h.authUsecase.SignIn(reqBody)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	cfg := config.GetConfig()

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(time.Second * time.Duration(cfg.Jwt.Expiration)),
		HTTPOnly: true,
	})

	response := NewResponse(true, "Sign in successfully", nil)
	return SendResponse(c, response)
}
