package repositories

import "github.com/wiraphatys/shop-management-go/database"

type ProductRepository interface {
	FindAllProducts() (*[]database.Product, error)
}
