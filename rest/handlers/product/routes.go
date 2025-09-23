package product

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *TSNewHandler) RegisterRoutes(mux *http.ServeMux, chain *middleware.TSManager) {
	mux.Handle("GET /test",
		chain.With(http.HandlerFunc(h.Test)))
	mux.Handle("GET /route",
		chain.With(http.HandlerFunc(h.Test)))
	mux.Handle("GET /products",
		chain.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("POST /products",
		chain.With(http.HandlerFunc(h.CreateProduct), h.middlewares.RequireAuth))
	mux.Handle("GET /products/{id}",
		chain.With(http.HandlerFunc(h.Product)))
	mux.Handle("PUT /products/{id}", chain.With(http.HandlerFunc(h.SetProduct), h.middlewares.RequireAuth))
	mux.Handle("DELETE /products/{id}", chain.With(http.HandlerFunc(h.DeleteProduct), h.middlewares.RequireAuth))
}
