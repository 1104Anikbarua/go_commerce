package product

import (
	// "ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

func (h *TSNewHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.product.List()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	utils.SendData(w, products, http.StatusOK)

}
