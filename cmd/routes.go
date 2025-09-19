package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func initRoute(mux *http.ServeMux, chain *middleware.TSManager) {
	mux.Handle("GET /test",
		chain.With(http.HandlerFunc(handlers.Test)))
	mux.Handle("GET /route",
		chain.With(http.HandlerFunc(handlers.Test)))
	mux.Handle("GET /products",
		chain.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products",
		chain.With(http.HandlerFunc(handlers.AddProducts)))
	mux.Handle("GET /products/{id}",
		chain.With(http.HandlerFunc(handlers.Product)))
}
