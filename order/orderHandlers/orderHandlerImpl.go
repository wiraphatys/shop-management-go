package orderHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/shop-management-go/order/orderEntities"
	"github.com/wiraphatys/shop-management-go/order/orderUsecases"
)

type orderHandlerImpl struct {
	orderUsecase orderUsecases.OrderUsecase
}

func NewOrderHandler(orderUsecase orderUsecases.OrderUsecase) OrderHandler {
	return &orderHandlerImpl{
		orderUsecase: orderUsecase,
	}
}

func (h *orderHandlerImpl) CreateOrder(c *fiber.Ctx) error {
	reqBody := new(orderEntities.OrderData)
	if err := c.BodyParser(reqBody); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	result, err := h.orderUsecase.CreateOrder(reqBody)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Create order successful", result)
	return SendResponse(c, response)
}
