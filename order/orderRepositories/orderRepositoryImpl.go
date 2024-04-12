package orderRepositories

import (
	"fmt"

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

func (r *orderRepositoryImpl) FindAllOrders() (*[]database.Order, error) {
	var orders []database.Order
	result := r.db.Preload("OrderLines").Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return &orders, nil
}

func (r *orderRepositoryImpl) FindOrderById(o_id string) (*database.Order, error) {
	var order database.Order
	result := r.db.Preload("OrderLines").First(&order, "o_id = ?", o_id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Errorf("FindOrderById: %v", result.Error)
		return nil, result.Error
	}
	return &order, nil
}

func (r *orderRepositoryImpl) CreateOrder(c_id string, orderLines *[]database.OrderLine) (*database.Order, error) {
	// define repository of product , customer
	productRepository := productRepositories.NewProductRepository(r.db)

	// begin transaction
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Errorf("CreateOrder: Recovered from panic: %v", r)
		}
	}()

	// validate p_id
	for _, line := range *orderLines {
		// product existed or not ?
		if _, err := productRepository.FindProductById(line.PID); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("invalid p_id: %v", line.PID)
		}
	}

	// create new order
	order, err := r.InsertOrder(c_id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

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

func (r *orderRepositoryImpl) UpdateOrderLineById(orderLine *database.OrderLine) (*database.OrderLine, error) {
	// find order_lines by p_id and o_id
	result := r.db.Model(&orderLine).Updates(orderLine)
	if result.Error != nil {
		return nil, result.Error
	}
	return orderLine, nil
}
