package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
)

func Server() {
	cnf := config.GetConfig()
	userHandler := user.NewHandler()
	productHandler := product.NewHandler()
	reviewHandler := review.NewHandler()
	server := rest.NewServer(cnf, userHandler, productHandler, reviewHandler)
	server.Start()
}
