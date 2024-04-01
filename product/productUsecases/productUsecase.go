package productUsecases

import (
	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/product/productEntities"
)

type ProductUsecase interface {
	GetAllProducts() (*[]database.Product, error)
	GetProductById(p_id string) (*database.Product, error)
	CreateProduct(product *database.Product) (*database.Product, error)
	UpdateProductById(p_id string, productData *productEntities.ProductData) (*database.Product, error)
	DeleteProductById(p_id string) error
}
