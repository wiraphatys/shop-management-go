package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wiraphatys/shop-management-go/config"
	"github.com/wiraphatys/shop-management-go/customer/handlers"
	"github.com/wiraphatys/shop-management-go/customer/repositories"
	"github.com/wiraphatys/shop-management-go/customer/usecases"
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

	s.initializeCustomerHttpHandler()

	log.Printf("Server is starting on %v", url)
	if err := s.app.Listen(url); err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}

func (s *fiberServer) initializeCustomerHttpHandler() {
	// initialize all layer
	customerRepository := repositories.NewCustomerRepository(s.db)
	customerUsecase := usecases.NewCustomerUsecase(customerRepository)
	customerHandler := handlers.NewCustomerHandler(customerUsecase)

	// route
	customerRouter := s.app.Group("/api/v1/customer")
	customerRouter.Get("/", customerHandler.GetAllCustomers)
	customerRouter.Get("/:pid", customerHandler.GetCustomerByEmail)
	customerRouter.Post("/register", customerHandler.RegisterCustomer)
}
