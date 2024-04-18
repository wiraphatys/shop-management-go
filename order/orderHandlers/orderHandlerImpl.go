package orderHandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/shop-management-go/database"
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

func (h *orderHandlerImpl) GetAllOrders(c *fiber.Ctx) error {
	result, err := h.orderUsecase.GetAllOrders()
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Get all orders successful", result)
	return SendResponse(c, response)
}

func (h *orderHandlerImpl) GetOrderById(c *fiber.Ctx) error {
	o_id := strings.Trim(c.Params("o_id"), " ")

	result, err := h.orderUsecase.GetOrderById(o_id)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Get order successful", result)
	return SendResponse(c, response)
}

func (h *orderHandlerImpl) CreateOrder(c *fiber.Ctx) error {
	orderData := new(orderEntities.OrderData)
	if err := c.BodyParser(orderData); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	// validate orderData
	if orderData.CID == "" || orderData.OrderLines == nil || len(orderData.OrderLines) == 0 {
		response := NewResponse(false, "cid or order_lines is null", nil)
		return SendResponse(c, response)
	}

	result, err := h.orderUsecase.CreateOrder(orderData)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Create order successful", result)
	return SendResponse(c, response)
}

func (h *orderHandlerImpl) UpdateOrderLineById(c *fiber.Ctx) error {
	orderLine := new(database.OrderLine)
	if err := c.BodyParser(orderLine); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	// validate o_id , p_id
	if orderLine.OID == "" || orderLine.PID == "" {
		response := NewResponse(false, "o_id or p_id is null", nil)
		return SendResponse(c, response)
	}

	result, err := h.orderUsecase.UpdateOrderLineById(orderLine)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Update order_line successful", result)
	return SendResponse(c, response)
}

func (h *orderHandlerImpl) DeleteOrderById(c *fiber.Ctx) error {
	o_id := strings.Trim(c.Params("o_id"), " ")

	err := h.orderUsecase.DeleteOrderById(o_id)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Delete order successful", nil)
	return SendResponse(c, response)
}

func (h *orderHandlerImpl) DeleteOrderLineById(c *fiber.Ctx) error {
	orderLine := new(orderEntities.OrderLineData)
	if err := c.BodyParser(orderLine); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	// validate o_id , p_id
	if orderLine.OID == "" || orderLine.PID == "" {
		response := NewResponse(false, "o_id or p_id is null", nil)
		return SendResponse(c, response)
	}

	err := h.orderUsecase.DeleteOrderLineById(orderLine)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Delete order_line successful", nil)
	return SendResponse(c, response)
}
