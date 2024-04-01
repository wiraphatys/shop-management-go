package productRepositories

import (
	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/product/productEntities"
)

type ProductRepository interface {
	FindAllProducts() (*[]database.Product, error)
	FindProductById(p_id string) (*database.Product, error)
	InsertProduct(product *database.Product) (*database.Product, error)
	UpdateProductById(p_id string, productData *productEntities.ProductData) (*database.Product, error)
}
