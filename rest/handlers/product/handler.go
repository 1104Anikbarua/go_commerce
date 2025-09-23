package product

import (
	"ecommerce/rest/middlewares"
)

type TSNewHandler struct {
	middlewares middlewares.TSMiddleware
}

func NewHandler(middlewares middlewares.TSMiddleware) *TSNewHandler {
	return &TSNewHandler{
		middlewares: middlewares,
	}
}
