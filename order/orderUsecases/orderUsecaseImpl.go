package orderUsecases

import (
	"fmt"

	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/order/orderEntities"
	"github.com/wiraphatys/shop-management-go/order/orderRepositories"
)

type orderUsecaseImpl struct {
	orderRepository orderRepositories.OrderRepositoriy
}

func NewOrderUsecase(orderRepository orderRepositories.OrderRepositoriy) OrderUsecase {
	return &orderUsecaseImpl{
		orderRepository: orderRepository,
	}
}

func (u *orderUsecaseImpl) GetAllOrders() (*[]database.Order, error) {
	orders, err := u.orderRepository.FindAllOrders()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (u *orderUsecaseImpl) GetOrderById(o_id string) (*database.Order, error) {
	order, err := u.orderRepository.FindOrderById(o_id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (u *orderUsecaseImpl) CreateOrder(orderData *orderEntities.OrderData) (*database.Order, error) {
	// validate orderData
	if orderData.CID == "" || orderData.OrderLines == nil || len(orderData.OrderLines) == 0 {
		return nil, fmt.Errorf("cid or order_lines is null")
	}

	for idx, line := range orderData.OrderLines {
		if line.PID == "" || line.Quantity <= 0 {
			return nil, fmt.Errorf("missing PID or Quantity at index %d", idx)
		}
	}

	order, err := u.orderRepository.CreateOrder(orderData.CID, &orderData.OrderLines)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (u *orderUsecaseImpl) UpdateOrderLineById(orderLine *database.OrderLine) (*database.OrderLine, error) {
	// validate o_id , p_id
	if orderLine.OID == "" || orderLine.PID == "" {
		return nil, fmt.Errorf("o_id or p_id is null")
	}

	// validate quantity
	if orderLine.Quantity <= 0 {
		return nil, fmt.Errorf("quantity must be positive")
	}

	updatedOrderLine, err := u.orderRepository.UpdateOrderLineById(orderLine)
	if err != nil {
		return nil, err
	}

	return updatedOrderLine, nil
}

func (u *orderUsecaseImpl) DeleteOrderById(o_id string) error {
	if err := u.orderRepository.DeleteOrderById(o_id); err != nil {
		return err
	}
	return nil
}

func (u *orderUsecaseImpl) DeleteOrderLineById(orderLine *orderEntities.OrderLineData) error {
	// validate o_id , p_id
	if orderLine.OID == "" || orderLine.PID == "" {
		return fmt.Errorf("o_id or p_id is null")
	}

	if err := u.orderRepository.DeleteOrderLineById(orderLine.OID, orderLine.PID); err != nil {
		return err
	}

	return nil
}
