package product

import (
	// "ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *TSNewHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var productIdStr = r.PathValue("id")
	var productId, err = strconv.Atoi(productIdStr)
	if err != nil {
		http.Error(w, "Fail to convert id from string to int", http.StatusBadRequest)
		return
	}
	// database.DeleteProduct(productId)
	err = h.product.Delete(productId)
	if err != nil {
		http.Error(w, "Fail to convert id from string to int", http.StatusBadRequest)
		return
	}
	utils.SendData(w, "Data Deleted Successfully", http.StatusAccepted)
}
