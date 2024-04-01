package customerRepositories

import (
	"github.com/wiraphatys/shop-management-go/customer/customerEntities"
	"github.com/wiraphatys/shop-management-go/database"
)

type CustomerRepository interface {
	FindAllCustomers() (*[]database.Customer, error)
	FindCustomerByEmail(email string) (*database.Customer, error)
	InsertCustomer(customer *database.Customer) (*database.Customer, error)
	UpdateCustomerByEmail(email string, customerData *customerEntities.CustomerData) (*database.Customer, error)
	DeleteCustomerByEmail(email string) error
}
