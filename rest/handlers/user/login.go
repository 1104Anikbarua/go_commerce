package user

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type TSPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *TSNewHandler) Login(w http.ResponseWriter, r *http.Request) {

	var payload TSPayload

	var decoder = json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "Please Send Valid Json", http.StatusNotAcceptable)
		return
	}

	isUserOk := database.Find(payload.Email, payload.Password)
	if isUserOk == nil {
		http.Error(w, "Please Check Email and Password", http.StatusUnauthorized)
		return
	}
	cnf := config.GetConfig()
	accessToken, err := utils.CreateJwt(utils.TSPayload{
		Sub:         isUserOk.ID,
		FirstName:   isUserOk.FirstName,
		LastName:    isUserOk.LastName,
		IsShopOwner: false,
		Email:       isUserOk.Email,
	}, cnf.JwtSecretKey)
	if err != nil {
		http.Error(w, "Please login again", http.StatusUnauthorized)
		return
	}
	utils.SendData(w, accessToken, http.StatusOK)
}
