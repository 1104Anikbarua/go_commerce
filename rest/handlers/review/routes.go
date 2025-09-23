package review

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *TSNewHandler) RegisterRoutes(mux *http.ServeMux, chain *middleware.TSManager) {
	mux.Handle("GET /reviews", chain.With(http.HandlerFunc(h.GetReviews)))
}
