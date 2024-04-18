package customerUsecases

import (
	"github.com/wiraphatys/shop-management-go/customer/customerEntities"
	"github.com/wiraphatys/shop-management-go/customer/customerRepositories"
	"github.com/wiraphatys/shop-management-go/database"
)

type customerUsecaseImpl struct {
	customerRepository customerRepositories.CustomerRepository
}

func NewCustomerUsecase(customerRepository customerRepositories.CustomerRepository) CustomerUsecase {
	return &customerUsecaseImpl{
		customerRepository: customerRepository,
	}
}

func (u *customerUsecaseImpl) GetAllCustomers() (*[]database.Customer, error) {
	result, err := u.customerRepository.FindAllCustomers()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (u *customerUsecaseImpl) GetCustomerByEmail(email string) (*database.Customer, error) {
	result, err := u.customerRepository.FindCustomerByEmail(email)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *customerUsecaseImpl) RegisterCustomer(customer *database.Customer) (*database.Customer, error) {
	result, err := u.customerRepository.InsertCustomer(customer)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *customerUsecaseImpl) UpdateCustomerByEmail(email string, reqBody *customerEntities.CustomerData) (*database.Customer, error) {
	result, err := u.customerRepository.UpdateCustomerByEmail(email, reqBody)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *customerUsecaseImpl) DeleteCustomerByEmail(email string) error {
	if err := u.customerRepository.DeleteCustomerByEmail(email); err != nil {
		return err
	}

	return nil
}
