package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/shop-management-go/config"
	"github.com/wiraphatys/shop-management-go/customer/customerHandlers"
	"github.com/wiraphatys/shop-management-go/customer/customerRepositories"
	"github.com/wiraphatys/shop-management-go/customer/customerUsecases"
	"github.com/wiraphatys/shop-management-go/order/orderHandlers"
	"github.com/wiraphatys/shop-management-go/order/orderRepositories"
	"github.com/wiraphatys/shop-management-go/order/orderUsecases"
	"github.com/wiraphatys/shop-management-go/product/productHandlers"
	"github.com/wiraphatys/shop-management-go/product/productRepositories"
	"github.com/wiraphatys/shop-management-go/product/productUsecases"

	"gorm.io/gorm"
)

type fiberServer struct {
	app *fiber.App
	db  *gorm.DB
	cfg *config.Config
}

func NewFiberServer(cfg *config.Config, db *gorm.DB) Server {
	return &fiberServer{
		app: fiber.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *fiberServer) Start() {
	url := fmt.Sprintf("%v:%d", s.cfg.Server.Host, s.cfg.Server.Port)

	// init module
	s.initializeCustomerHttpHandler()
	s.initializeProductHttpHandler()
	s.initializeOrderHttpHandler()

	log.Printf("Server is starting on %v", url)
	if err := s.app.Listen(url); err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}

func (s *fiberServer) initializeCustomerHttpHandler() {
	// initialize all layer
	customerRepository := customerRepositories.NewCustomerRepository(s.db)
	customerUsecase := customerUsecases.NewCustomerUsecase(customerRepository)
	customerHandler := customerHandlers.NewCustomerHandler(customerUsecase)

	// route
	customerRouter := s.app.Group("/api/v1/customer")
	customerRouter.Get("/", customerHandler.GetAllCustomers)
	customerRouter.Get("/:email", customerHandler.GetCustomerByEmail)
	customerRouter.Post("/register", customerHandler.RegisterCustomer)
	customerRouter.Put("/:email", customerHandler.UpdateCustomerByEmail)
	customerRouter.Delete("/:email", customerHandler.DeleteCustomerByEmail)
}

func (s *fiberServer) initializeProductHttpHandler() {
	// initialize all layer
	productRepository := productRepositories.NewProductRepository(s.db)
	productUsecase := productUsecases.NewProductUsecase(productRepository)
	productHandler := productHandlers.NewProductHandler(productUsecase)

	// route
	productRouter := s.app.Group("/api/v1/product")
	productRouter.Get("/", productHandler.GetAllProducts)
	productRouter.Get("/:p_id", productHandler.GetProductById)
	productRouter.Post("/", productHandler.CreateProduct)
	productRouter.Put("/:p_id", productHandler.UpdateProductById)
	productRouter.Delete("/:p_id", productHandler.DeleteProductById)
}

func (s *fiberServer) initializeOrderHttpHandler() {
	// initialize all layer
	orderRepository := orderRepositories.NewOrderRepository(s.db)
	orderUsecase := orderUsecases.NewOrderUsecase(orderRepository)
	orderHandler := orderHandlers.NewOrderHandler(orderUsecase)

	// route
	orderRouter := s.app.Group("/api/v1/order")
	orderRouter.Get("/", orderHandler.GetAllOrders)
	orderRouter.Get("/:o_id", orderHandler.GetOrderById)
	orderRouter.Post("/", orderHandler.CreateOrder)
	orderRouter.Put("/", orderHandler.UpdateOrderLineById)
	orderRouter.Delete("/:o_id", orderHandler.DeleteOrderById)
	orderRouter.Delete("/", orderHandler.DeleteOrderLineById)
}
