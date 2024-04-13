package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/shop-management-go/admin/adminHandlers"
	"github.com/wiraphatys/shop-management-go/admin/adminRepositories"
	"github.com/wiraphatys/shop-management-go/admin/adminUsecases"
	"github.com/wiraphatys/shop-management-go/auth/authHandlers"
	"github.com/wiraphatys/shop-management-go/auth/authUsecases"
	"github.com/wiraphatys/shop-management-go/config"
	"github.com/wiraphatys/shop-management-go/customer/customerHandlers"
	"github.com/wiraphatys/shop-management-go/customer/customerRepositories"
	"github.com/wiraphatys/shop-management-go/customer/customerUsecases"
	"github.com/wiraphatys/shop-management-go/middleware/middlewareHandlers"
	"github.com/wiraphatys/shop-management-go/middleware/middlewareUsecases"
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

	// init middleware
	middleware := s.initializeMiddlewareHttpHandler()

	// init module
	s.initializeAuthHttpHandler(middleware)
	s.initializeCustomerHttpHandler(middleware)
	s.initializeProductHttpHandler(middleware)
	s.initializeOrderHttpHandler(middleware)
	s.initializeAdminHttpHandler(middleware)

	log.Printf("Server is starting on %v", url)
	if err := s.app.Listen(url); err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}

func (s *fiberServer) initializeMiddlewareHttpHandler() middlewareHandlers.MiddlewareHandler {
	// initialize all layer
	adminRepository := adminRepositories.NewAdminRepository(s.db)
	middlewareUsecase := middlewareUsecases.NewMiddlewareUsecase(adminRepository)
	middlewareHandler := middlewareHandlers.NewMiddlewareHandler(middlewareUsecase)

	return middlewareHandler
}

func (s *fiberServer) initializeAuthHttpHandler(middleware middlewareHandlers.MiddlewareHandler) {
	// initialize all layer
	adminRepository := adminRepositories.NewAdminRepository(s.db)
	authUsecase := authUsecases.NewAuthUsecase(adminRepository)
	authHandler := authHandlers.NewAuthHandler(authUsecase)

	// route
	authRouter := s.app.Group("/api/v1/auth")
	authRouter.Post("/signin", authHandler.SignIn)
	authRouter.Post("/signout", middleware.Authenticated, authHandler.SignOut)
}

func (s *fiberServer) initializeCustomerHttpHandler(middleware middlewareHandlers.MiddlewareHandler) {
	// initialize all layer
	customerRepository := customerRepositories.NewCustomerRepository(s.db)
	customerUsecase := customerUsecases.NewCustomerUsecase(customerRepository)
	customerHandler := customerHandlers.NewCustomerHandler(customerUsecase)

	// route
	customerRouter := s.app.Group("/api/v1/customer")
	customerRouter.Get("/", middleware.Authenticated, customerHandler.GetAllCustomers)
	customerRouter.Get("/:email", middleware.Authenticated, customerHandler.GetCustomerByEmail)
	customerRouter.Post("/register", middleware.Authenticated, customerHandler.RegisterCustomer)
	customerRouter.Put("/:email", middleware.Authenticated, customerHandler.UpdateCustomerByEmail)
	customerRouter.Delete("/:email", middleware.Authenticated, customerHandler.DeleteCustomerByEmail)
}

func (s *fiberServer) initializeProductHttpHandler(middleware middlewareHandlers.MiddlewareHandler) {
	// initialize all layer
	productRepository := productRepositories.NewProductRepository(s.db)
	productUsecase := productUsecases.NewProductUsecase(productRepository)
	productHandler := productHandlers.NewProductHandler(productUsecase)

	// route
	productRouter := s.app.Group("/api/v1/product")
	productRouter.Get("/", middleware.Authenticated, productHandler.GetAllProducts)
	productRouter.Get("/:p_id", middleware.Authenticated, productHandler.GetProductById)
	productRouter.Post("/", middleware.Authenticated, productHandler.CreateProduct)
	productRouter.Put("/:p_id", middleware.Authenticated, productHandler.UpdateProductById)
	productRouter.Delete("/:p_id", middleware.Authenticated, productHandler.DeleteProductById)
}

func (s *fiberServer) initializeOrderHttpHandler(middleware middlewareHandlers.MiddlewareHandler) {
	// initialize all layer
	orderRepository := orderRepositories.NewOrderRepository(s.db)
	orderUsecase := orderUsecases.NewOrderUsecase(orderRepository)
	orderHandler := orderHandlers.NewOrderHandler(orderUsecase)

	// route
	orderRouter := s.app.Group("/api/v1/order")
	orderRouter.Get("/", middleware.Authenticated, orderHandler.GetAllOrders)
	orderRouter.Get("/:o_id", middleware.Authenticated, orderHandler.GetOrderById)
	orderRouter.Post("/", middleware.Authenticated, orderHandler.CreateOrder)
	orderRouter.Put("/", middleware.Authenticated, orderHandler.UpdateOrderLineById)
	orderRouter.Delete("/:o_id", middleware.Authenticated, orderHandler.DeleteOrderById)
	orderRouter.Delete("/", middleware.Authenticated, orderHandler.DeleteOrderLineById)
}

func (s *fiberServer) initializeAdminHttpHandler(middleware middlewareHandlers.MiddlewareHandler) {
	// initialize all layer
	adminRepository := adminRepositories.NewAdminRepository(s.db)
	adminUsecase := adminUsecases.NewAdminUsecase(adminRepository)
	adminHandler := adminHandlers.NewAdminHandler(adminUsecase)

	// route
	adminRouter := s.app.Group("/api/v1/admin")
	adminRouter.Get("/:email", middleware.Authenticated, adminHandler.GetAdminByEmail)
	adminRouter.Post("/", middleware.Authenticated, adminHandler.CreateAdmin)
}
