package usecases

import "github.com/wiraphatys/shop-management-go/database"

type CustomerUsecase interface {
	GetAllCustomers() (*[]database.Customer, error)
	GetCustomerByEmail(email string) (*database.Customer, error)
	RegisterCustomer(customer *database.Customer) (*database.Customer, error)
	DeleteCustomerByEmail(email string) error
}
