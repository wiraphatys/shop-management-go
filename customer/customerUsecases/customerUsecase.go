package customerUsecases

import (
	"github.com/wiraphatys/shop-management-go/customer/customerEntities"
	"github.com/wiraphatys/shop-management-go/database"
)

type CustomerUsecase interface {
	GetAllCustomers() (*[]database.Customer, error)
	GetCustomerByEmail(email string) (*database.Customer, error)
	RegisterCustomer(customer *database.Customer) (*database.Customer, error)
	UpdateCustomerByEmail(email string, reqBody *customerEntities.CustomerData) (*database.Customer, error)
	DeleteCustomerByEmail(email string) error
}
