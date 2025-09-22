package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser database.TSUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Please Send Valid Json", http.StatusNotAcceptable)
		return
	}
	user := newUser.Store()

	utils.SendData(w, user, http.StatusCreated)
}
