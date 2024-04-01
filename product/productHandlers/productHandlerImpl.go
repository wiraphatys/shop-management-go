package productHandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/product/productEntities"
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

func (h *productHandlerImpl) GetProductById(c *fiber.Ctx) error {
	p_id := strings.Trim(c.Params("p_id"), " ")

	result, err := h.productUsecase.GetProductById(p_id)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Get product successful", result)
	return SendResponse(c, response)
}

func (h *productHandlerImpl) CreateProduct(c *fiber.Ctx) error {
	reqBody := new(database.Product)
	if err := c.BodyParser(reqBody); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	product, err := h.productUsecase.CreateProduct(reqBody)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Create product successful", product)
	return SendResponse(c, response)
}

func (h *productHandlerImpl) UpdateProductById(c *fiber.Ctx) error {
	reqBody := new(productEntities.ProductData)
	if err := c.BodyParser(reqBody); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	p_id := strings.Trim(c.Params("p_id"), " ")

	result, err := h.productUsecase.UpdateProductById(p_id, reqBody)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Update product successful", result)
	return SendResponse(c, response)
}

func (h *productHandlerImpl) DeleteProductById(c *fiber.Ctx) error {
	p_id := strings.Trim(c.Params("p_id"), " ")

	if err := h.productUsecase.DeleteProductById(p_id); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Delete product successful", nil)
	return SendResponse(c, response)
}
