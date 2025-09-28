package product

import (
	// "ecommerce/database"
	"ecommerce/repository"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *TSNewHandler) SetProduct(w http.ResponseWriter, r *http.Request) {

	productIdStr := r.PathValue("id")
	id, err := strconv.Atoi(productIdStr)

	if err != nil {
		http.Error(w, "Please Send Valid Json", http.StatusBadRequest)
		return
	}
	var payload repository.TSProducts //database.TSProducts
	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&payload)
	if error != nil {
		http.Error(w, "Please Send Valid Json", http.StatusBadRequest)
		return
	}
	payload.Id = id

	product, err := h.product.Update(payload)
	if err != nil {
		http.Error(w, "Fail To Update Product", http.StatusNotImplemented)
	}
	// database.SetProduct(payload)
	utils.SendData(w, product, http.StatusAccepted)
}
