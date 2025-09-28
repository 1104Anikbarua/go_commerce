package product

import (
	"ecommerce/repository"
	"ecommerce/rest/middlewares"
)

type TSNewHandler struct {
	middlewares *middlewares.TSMiddleware
	product     repository.TIProductRepo
}

func NewHandler(middlewares *middlewares.TSMiddleware, product repository.TIProductRepo) *TSNewHandler {
	return &TSNewHandler{
		middlewares: middlewares,
		product:     product,
	}
}
