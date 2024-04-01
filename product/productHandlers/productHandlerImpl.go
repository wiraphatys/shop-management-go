package productHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/shop-management-go/product/productUsecases"
)

type productHandlerImpl struct {
	productUsecase productUsecases.ProductUsecase
}

func NewProductHandler(productUsecase productUsecases.ProductUsecase) ProductHandler {
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
