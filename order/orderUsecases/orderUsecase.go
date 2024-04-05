package orderUsecases

import (
	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/order/orderEntities"
)

type OrderUsecase interface {
	CreateOrder(orderData *orderEntities.OrderData) (*database.Order, error)
}
