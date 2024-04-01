package productRepositories

import (
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/product/productEntities"
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
	var products []database.Product
	result := r.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
}
func (r *productRepositoryImpl) FindProductById(p_id string) (*database.Product, error) {
	var product database.Product
	result := r.db.First(&product, "p_id = ?", p_id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Errorf("FindProductById: %v", result.Error)
		return nil, result.Error
	}
	return &product, nil
}

func (r *productRepositoryImpl) InsertProduct(product *database.Product) (*database.Product, error) {
	// create new product
	product.PID = "1"
	result := r.db.Create(product)
	if result.Error != nil {
		log.Errorf("InsertProduct: %v", result.Error)
		return nil, result.Error
	}
	log.Debugf("InsertProduct: %v", result.RowsAffected)

	// Get the last inserted product
	var createdProduct database.Product
	err := r.db.Order("p_id desc").First(&createdProduct).Error
	if err != nil {
		log.Errorf("GetLastProduct: %v", err)
		return nil, err
	}

	return &createdProduct, nil
}

func (r *productRepositoryImpl) UpdateProductById(p_id string, productData *productEntities.ProductData) (*database.Product, error) {
	var product *database.Product

	// Find product by id
	result := r.db.Where("p_id = ?", p_id).First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, result.Error
	}

	// update the product fields
	updateMap := make(map[string]interface{})
	if productData.Name != "" {
		updateMap["name"] = productData.Name
	}
	if productData.Description != "" {
		updateMap["description"] = productData.Description
	}
	if productData.UnitPrice != 0 {
		updateMap["unit_price"] = productData.UnitPrice
	}

	// save product to database
	result = r.db.Model(&product).Updates(updateMap)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *productRepositoryImpl) DeleteProductById(p_id string) error {
	result := r.db.Unscoped().Where("p_id = ?", p_id).Delete(&database.Product{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}
