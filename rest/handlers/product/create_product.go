package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func (h *TSNewHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.TSProducts

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Please Send Valid Json", http.StatusBadRequest)
		return
	}
	// database.Products = append(database.Products, newProduct)
	product := database.Store(newProduct)
	utils.SendData(w, product, http.StatusOK)

}
