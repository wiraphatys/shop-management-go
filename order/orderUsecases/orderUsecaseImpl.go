package orderUsecases

import (
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

func (u *orderUsecaseImpl) CreateOrder(orderData *orderEntities.OrderData) (*database.Order, error) {
	order, err := u.orderRepository.CreateOrder(orderData.CID, &orderData.OrderLines)
	if err != nil {
		return nil, err
	}

	return order, nil
}
