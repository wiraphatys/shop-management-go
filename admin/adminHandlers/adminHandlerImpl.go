package adminHandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/shop-management-go/admin/adminEntities"
	"github.com/wiraphatys/shop-management-go/admin/adminUsecases"
	"github.com/wiraphatys/shop-management-go/database"
)

type adminHandlerImpl struct {
	adminUsecase adminUsecases.AdminUsecase
}

func NewAdminHandler(adminUsecase adminUsecases.AdminUsecase) AdminHandler {
	return &adminHandlerImpl{
		adminUsecase: adminUsecase,
	}
}

func (h *adminHandlerImpl) GetAdminByEmail(c *fiber.Ctx) error {
	email := strings.Trim(c.Params("email"), " ")

	admin, err := h.adminUsecase.GetAdminByEmail(email)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Get admin successful", admin)
	return SendResponse(c, response)
}

func (h *adminHandlerImpl) CreateAdmin(c *fiber.Ctx) error {
	reqBody := new(database.Admin)
	if err := c.BodyParser(reqBody); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	admin, err := h.adminUsecase.CreateAdmin(reqBody)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	adminResponse := adminEntities.AdminResponse{
		AID:       admin.AID,
		Email:     admin.Email,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}

	response := NewResponse(true, "Register admin successful", adminResponse)
	return SendResponse(c, response)
}
