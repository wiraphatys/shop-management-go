package usecases

import (
	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/product/repositories"
)

type productUsecaseImpl struct {
	productRepository repositories.ProductRepository
}

func NewProductUsecase(productRepository repositories.ProductRepository) ProductUsecase {
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
