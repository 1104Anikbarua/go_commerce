package user

import (
	// "ecommerce/database"
	"ecommerce/repository"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type TSUser struct {
	ID           int    `json:"_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsShopOwnder bool   `json:"is_shop_owner"`
}

func (h *TSNewHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var req TSUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Please Send Valid Json", http.StatusNotAcceptable)
		return
	}
	// user := req.Store()
	user, err := h.userRepo.Create(repository.TSUser{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		Password:     req.Password,
		IsShopOwnder: req.IsShopOwnder,
	})

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	utils.SendData(w, user, http.StatusCreated)
}
