package repositories

import (
	"github.com/wiraphatys/shop-management-go/database"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{
		db: db,
	}
}

func (r *productRepositoryImpl) FindAllProducts() (*[]database.Product, error) {
	var products *[]database.Product
	result := r.db.Find(products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
