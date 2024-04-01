package repositories

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/wiraphatys/shop-management-go/database"
	"gorm.io/gorm"
)

type customerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepositoryImpl{
		db: db,
	}
}

func (r *customerRepositoryImpl) FindAllCustomers() (*[]database.Customer, error) {
	var customers []database.Customer
	result := r.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customers, nil
}

func (r *customerRepositoryImpl) FindCustomerByEmail(email string) (*database.Customer, error) {
	var customer database.Customer
	result := r.db.First(&customer, "email = ?", email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Errorf("FindUserByEmail: %v", result.Error)
		return nil, result.Error
	}
	return &customer, nil
}

func (r *customerRepositoryImpl) InsertCustomer(customer *database.Customer) (*database.Customer, error) {
	// create new customer
	customer.CID = "1"
	result := r.db.Create(customer)
	if result.Error != nil {
		log.Errorf("InsertUserData: %v", result.Error)
		return nil, result.Error
	}
	log.Debugf("InsertUserData: %v", result.RowsAffected)

	// returning created customer
	createdCustomer, err := r.FindCustomerByEmail(customer.Email)
	if err != nil {
		return nil, err
	}
	return createdCustomer, nil
}

func (r *customerRepositoryImpl) DeleteCustomerByEmail(email string) error {
	return nil
}
