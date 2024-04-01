package productUsecases

import (
	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/product/productRepositories"
)

type productUsecaseImpl struct {
	productRepository productRepositories.ProductRepository
}

func NewProductUsecase(productRepository productRepositories.ProductRepository) ProductUsecase {
	return &productUsecaseImpl{
		productRepository: productRepository,
	}
}

func (u *productUsecaseImpl) GetAllProducts() (*[]database.Product, error) {
	result, err := u.productRepository.FindAllProducts()
	if err != nil {
		return nil, err
	}
	return result, nil
}
