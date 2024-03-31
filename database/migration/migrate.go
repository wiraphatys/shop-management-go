package main

import (
	"fmt"

	"github.com/wiraphatys/shop-management-go/config"
	"github.com/wiraphatys/shop-management-go/database"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(cfg)

	// migrate schema
	db.GetDb().AutoMigrate(&database.Customer{})
	db.GetDb().AutoMigrate(&database.Product{})
	db.GetDb().AutoMigrate(&database.Admin{})
	db.GetDb().AutoMigrate(&database.Order{})
	db.GetDb().AutoMigrate(&database.OrderLine{})

	// migrate trigger functions
	if err := database.CreateCustomerIDTrigger(db.GetDb()); err != nil {
		panic(err)
	}

	if err := database.CreateProductIDTrigger(db.GetDb()); err != nil {
		panic(err)
	}

	if err := database.CreateOrderIDTrigger(db.GetDb()); err != nil {
		panic(err)
	}

	if err := database.CreateOrderLineIDTrigger(db.GetDb()); err != nil {
		panic(err)
	}

	if err := database.CreateAdminIDTrigger(db.GetDb()); err != nil {
		panic(err)
	}

	fmt.Printf("%v", "Migration Successful")
}
