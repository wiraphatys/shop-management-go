package productRepositories

import "github.com/wiraphatys/shop-management-go/database"

type ProductRepository interface {
	FindAllProducts() (*[]database.Product, error)
	FindProductById(p_id string) (*database.Product, error)
	InsertProduct(product *database.Product) (*database.Product, error)
}
