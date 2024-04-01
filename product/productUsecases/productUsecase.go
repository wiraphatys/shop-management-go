package productUsecases

import "github.com/wiraphatys/shop-management-go/database"

type ProductUsecase interface {
	GetAllProducts() (*[]database.Product, error)
}
