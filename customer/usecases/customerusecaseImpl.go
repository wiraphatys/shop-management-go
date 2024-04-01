package usecases

import (
	"fmt"

	"github.com/wiraphatys/shop-management-go/customer/entities"
	"github.com/wiraphatys/shop-management-go/customer/repositories"
	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/util"
)

type customerUsecaseImpl struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerUsecase(customerRepository repositories.CustomerRepository) CustomerUsecase {
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
	if !util.IsEmailValid(email) {
		return nil, fmt.Errorf("invalid email address")
	}

	result, err := u.customerRepository.FindCustomerByEmail(email)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (u *customerUsecaseImpl) RegisterCustomer(customer *database.Customer) (*database.Customer, error) {
	if !util.IsEmailValid(customer.Email) {
		return nil, fmt.Errorf("invalid email address")
	}

	result, err := u.customerRepository.InsertCustomer(customer)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *customerUsecaseImpl) UpdateCustomerByEmail(email string, reqBody *entities.CustomerData) (*database.Customer, error) {
	if !util.IsEmailValid(email) {
		return nil, fmt.Errorf("invalid email address")
	}

	result, err := u.customerRepository.UpdateCustomerByEmail(email, reqBody)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *customerUsecaseImpl) DeleteCustomerByEmail(email string) error {
	if !util.IsEmailValid(email) {
		return fmt.Errorf("invalid email address")
	}

	if err := u.customerRepository.DeleteCustomerByEmail(email); err != nil {
		return err
	}

	return nil
}
