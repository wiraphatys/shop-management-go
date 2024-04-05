package orderUsecases

import (
	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/order/orderEntities"
)

type OrderUsecase interface {
	GetAllOrders() (*[]database.Order, error)
	GetOrderById(o_id string) (*database.Order, error)
	CreateOrder(orderData *orderEntities.OrderData) (*database.Order, error)
}
