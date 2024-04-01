package repositories

import (
	"github.com/wiraphatys/shop-management-go/database"
)

type CustomerRepository interface {
	FindAllCustomers() (*[]database.Customer, error)
	FindCustomerByEmail(email string) (*database.Customer, error)
	InsertCustomer(customer *database.Customer) (*database.Customer, error)
	DeleteCustomerByEmail(email string) error
}
