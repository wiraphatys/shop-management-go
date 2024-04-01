package productUsecases

import (
	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/product/productEntities"
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

func (u *productUsecaseImpl) GetProductById(p_id string) (*database.Product, error) {
	result, err := u.productRepository.FindProductById(p_id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *productUsecaseImpl) CreateProduct(product *database.Product) (*database.Product, error) {
	result, err := u.productRepository.InsertProduct(product)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *productUsecaseImpl) UpdateProductById(p_id string, productData *productEntities.ProductData) (*database.Product, error) {
	result, err := u.productRepository.UpdateProductById(p_id, productData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
