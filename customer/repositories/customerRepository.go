package repositories

import (
	"github.com/wiraphatys/shop-management-go/customer/entities"
	"github.com/wiraphatys/shop-management-go/database"
)

type CustomerRepository interface {
	FindAllCustomers() (*[]database.Customer, error)
	FindCustomerByEmail(email string) (*database.Customer, error)
	InsertCustomer(customer *database.Customer) (*database.Customer, error)
	UpdateCustomerByEmail(email string, customerData *entities.CustomerData) (*database.Customer, error)
	DeleteCustomerByEmail(email string) error
}
