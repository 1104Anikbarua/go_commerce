package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func Server() {
	cnf := config.GetConfig()
	userHandler := user.NewHandler()
	middleware := middlewares.NewMiddleware(cnf)
	productHandler := product.NewHandler(middleware)
	reviewHandler := review.NewHandler()
	server := rest.NewServer(cnf, userHandler, productHandler, reviewHandler)
	server.Start()
}
