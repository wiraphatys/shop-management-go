package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/shop-management-go/customer/entities"
	"github.com/wiraphatys/shop-management-go/customer/usecases"
	"github.com/wiraphatys/shop-management-go/database"
)

type customerHandlerImpl struct {
	customerUsecase usecases.CustomerUsecase
}

func NewCustomerHandler(customerUsecase usecases.CustomerUsecase) CustomerHandler {
	return &customerHandlerImpl{
		customerUsecase: customerUsecase,
	}
}

func (h *customerHandlerImpl) GetAllCustomers(c *fiber.Ctx) error {
	result, err := h.customerUsecase.GetAllCustomers()
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Get all customer successful", result)
	return SendResponse(c, response)
}
func (h *customerHandlerImpl) GetCustomerByEmail(c *fiber.Ctx) error {
	email := strings.Trim(c.Params("email"), " ")

	result, err := h.customerUsecase.GetCustomerByEmail(email)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Get customer successful", result)
	return SendResponse(c, response)
}
func (h *customerHandlerImpl) RegisterCustomer(c *fiber.Ctx) error {
	reqBody := new(database.Customer)
	if err := c.BodyParser(reqBody); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	customer, err := h.customerUsecase.RegisterCustomer(reqBody)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Register customer successful", customer)
	return SendResponse(c, response)
}

func (h *customerHandlerImpl) UpdateCustomerByEmail(c *fiber.Ctx) error {
	reqBody := new(entities.CustomerData)
	if err := c.BodyParser(reqBody); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	email := strings.Trim(c.Params("email"), " ")

	result, err := h.customerUsecase.UpdateCustomerByEmail(email, reqBody)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Update customer successful", result)
	return SendResponse(c, response)
}

func (h *customerHandlerImpl) DeleteCustomerByEmail(c *fiber.Ctx) error {
	email := strings.Trim(c.Params("email"), " ")

	if err := h.customerUsecase.DeleteCustomerByEmail(email); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	response := NewResponse(true, "Delete customer successful", nil)
	return SendResponse(c, response)
}
