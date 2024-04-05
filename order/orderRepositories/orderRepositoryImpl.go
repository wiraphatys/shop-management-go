package orderRepositories

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/product/productRepositories"
	"gorm.io/gorm"
)

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepositoriy {
	return &orderRepositoryImpl{
		db: db,
	}
}

func (r *orderRepositoryImpl) CreateOrder(c_id string, orderLines *[]database.OrderLine) (*database.Order, error) {
	// begin transaction
	tx := r.db.Begin()
	defer func() {
		if e := recover(); e != nil {
			tx.Rollback()
		}
	}()

	// create new order
	order, err := r.InsertOrder(c_id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// define product repository
	productRepository := productRepositories.NewProductRepository(r.db)

	// create new orderlines
	for _, line := range *orderLines {
		// product existed or not ?
		if _, err := productRepository.FindProductById(line.PID); err != nil {
			tx.Rollback()
			return nil, err
		}

		// create orderlines
		if _, err := r.InsertOrderLine(order.OID, line.PID, line.Quantity); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// everything is right, commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var createdOrder database.Order
	if err := r.db.Preload("OrderLines").Where("o_id = ?", order.OID).First(&createdOrder).Error; err != nil {
		return nil, err
	}

	return &createdOrder, nil
}

func (r *orderRepositoryImpl) InsertOrder(c_id string) (*database.Order, error) {
	// prepare struct Order
	order := &database.Order{
		CID: c_id,
	}

	// create new order with c_id
	result := r.db.Create(order)
	if result.Error != nil {
		log.Errorf("InsertOrder: %v", result.Error)
		return nil, result.Error
	}

	// get latest order
	var createdOrder database.Order
	err := r.db.Order("o_id desc").First(&createdOrder).Error
	if err != nil {
		log.Errorf("GetLastOrder: %v", err)
		return nil, err
	}

	return &createdOrder, nil
}

func (r *orderRepositoryImpl) InsertOrderLine(o_id, p_id string, quantity int) (*database.OrderLine, error) {
	// prepare OrderLine struct
	orderLine := &database.OrderLine{
		OID:      o_id,
		PID:      p_id,
		Quantity: quantity,
	}

	// create new orderline
	result := r.db.Create(orderLine)
	if result.Error != nil {
		log.Errorf("InsertOrderLine: %v", result.Error)
		return nil, result.Error
	}

	return orderLine, nil
}
