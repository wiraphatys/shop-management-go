package productUsecases

import "github.com/wiraphatys/shop-management-go/database"

type ProductUsecase interface {
	GetAllProducts() (*[]database.Product, error)
	GetProductById(p_id string) (*database.Product, error)
	CreateProduct(product *database.Product) (*database.Product, error)
}
