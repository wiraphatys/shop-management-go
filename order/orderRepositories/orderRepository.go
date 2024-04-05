package orderRepositories

import (
	"github.com/wiraphatys/shop-management-go/database"
)

type OrderRepositoriy interface {
	CreateOrder(c_id string, orderLines *[]database.OrderLine) (*database.Order, error)
	InsertOrder(c_id string) (*database.Order, error)
	InsertOrderLine(o_id, p_id string, quantity int) (*database.OrderLine, error)
}
