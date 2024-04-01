package usecases

import (
	"github.com/wiraphatys/shop-management-go/customer/entities"
	"github.com/wiraphatys/shop-management-go/database"
)

type CustomerUsecase interface {
	GetAllCustomers() (*[]database.Customer, error)
	GetCustomerByEmail(email string) (*database.Customer, error)
	RegisterCustomer(customer *database.Customer) (*database.Customer, error)
	UpdateCustomerByEmail(email string, reqBody *entities.CustomerData) (*database.Customer, error)
	DeleteCustomerByEmail(email string) error
}
