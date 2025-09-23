package user

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *TSNewHandler) RegisterRoutes(mux *http.ServeMux, chain *middleware.TSManager) {
	mux.Handle("POST /users", chain.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("POST /login", chain.With(http.HandlerFunc(h.Login)))
}
