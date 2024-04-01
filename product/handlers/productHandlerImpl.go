package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/shop-management-go/product/usecases"
)

type productHandlerImpl struct {
	productUsecase usecases.ProductUsecase
}

func NewProductHandler(productUsecase usecases.ProductUsecase) ProductHandler {
	return &productHandlerImpl{
		productUsecase: productUsecase,
	}
}

func (h *productHandlerImpl) GetAllProducts(c *fiber.Ctx) error {
	result, err := h.productUsecase.GetAllProducts()
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Get all product successful", result)
	return SendResponse(c, response)
}
