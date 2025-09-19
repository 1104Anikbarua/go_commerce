package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func AddProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct database.TSProducts
	newProduct.Id = len(database.Products) + 1
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Please Send Valid Json", http.StatusBadRequest)
		return
	}
	database.Products = append(database.Products, newProduct)
	utils.SendData(w, newProduct, http.StatusOK)

}
