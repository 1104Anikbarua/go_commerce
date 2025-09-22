package rest

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
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
		chain.With(http.HandlerFunc(handlers.CreateProduct), middleware.RequireAuth))
	mux.Handle("GET /products/{id}",
		chain.With(http.HandlerFunc(handlers.Product)))
	mux.Handle("PUT /products/{id}", chain.With(http.HandlerFunc(handlers.SetProduct)))
	mux.Handle("DELETE /products/{id}", chain.With(http.HandlerFunc(handlers.DeleteProduct)))
	mux.Handle("POST /users", chain.With(http.HandlerFunc(handlers.CreateUser)))
	mux.Handle("POST /login", chain.With(http.HandlerFunc(handlers.Login)))
}
