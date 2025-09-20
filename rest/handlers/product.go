package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

func Product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w)
	fmt.Println("R--->", r)
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	fmt.Println("ID--->", id)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusNotFound)
	}
	// for _, product := range database.Products {
	// 	if product.Id == id {
	// 		utils.SendData(w, product, http.StatusOK)
	// 		return
	// 	}
	// }
	product := database.Product(id)

	if product == nil {
		utils.SendError(w, "Data Not Found", http.StatusNotFound)
		return
	}
	utils.SendData(w, product, http.StatusOK)
}
