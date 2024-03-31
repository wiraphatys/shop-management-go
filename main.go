package main

import (
	"github.com/wiraphatys/shop-management-go/config"
	"github.com/wiraphatys/shop-management-go/database"
	"github.com/wiraphatys/shop-management-go/server"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(cfg)

	server.NewFiberServer(cfg, db.GetDb()).Start()
}
