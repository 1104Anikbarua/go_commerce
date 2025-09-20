package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func SetProduct(w http.ResponseWriter, r *http.Request) {

	productIdStr := r.PathValue("id")
	id, err := strconv.Atoi(productIdStr)

	if err != nil {
		http.Error(w, "Please Send Valid Json", http.StatusBadRequest)
		return
	}
	var payload database.TSProducts
	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&payload)
	if error != nil {
		http.Error(w, "Please Send Valid Json", http.StatusBadRequest)
		return
	}
	payload.Id = id
	database.SetProduct(payload)
	utils.SendData(w, "Data Updated Successfully", http.StatusAccepted)
}
