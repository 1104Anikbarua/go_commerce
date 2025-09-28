package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type TSServer struct {
	cnf            *config.TSConfig
	userHandler    *user.TSNewHandler
	productHandler *product.TSNewHandler
	reviewHandler  *review.TSNewHandler
}

func NewServer(cnf *config.TSConfig, userHandler *user.TSNewHandler, productHandler *product.TSNewHandler, reviewHandler *review.TSNewHandler) *TSServer {
	return &TSServer{
		cnf:            cnf,
		userHandler:    userHandler,
		productHandler: productHandler,
		reviewHandler:  reviewHandler,
	}
}

func (server *TSServer) Start() {
	chain := middleware.NewManager()
	chain.Use(middleware.Preflight, middleware.Cors, middleware.Test, middleware.Logger)
	mux := http.NewServeMux()
	wrapedMux := chain.WrapMux(mux)

	server.userHandler.RegisterRoutes(mux, chain)
	server.productHandler.RegisterRoutes(mux, chain)
	server.reviewHandler.RegisterRoutes(mux, chain)

	port := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server Running On Port", port)
	err := http.ListenAndServe(port, wrapedMux)
	if err != nil {
		fmt.Println("Fail to run the server", err)
		os.Exit(1)
	}
}
