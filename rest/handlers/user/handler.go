package user

import (
	"ecommerce/repository"
)

type TSNewHandler struct {
	userRepo repository.TIUserRepo
}

func NewHandler() *TSNewHandler {
	return &TSNewHandler{}
}
