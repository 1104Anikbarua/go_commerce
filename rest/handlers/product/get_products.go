package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

func (h *TSNewHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.List(), http.StatusOK)

}
