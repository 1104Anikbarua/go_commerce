package product

import (
	// "ecommerce/database"
	// "ecommerce/database"
	"ecommerce/repository"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type TSReqProducts struct {
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *TSNewHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req TSReqProducts //database.TSProducts

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Please Send Valid Json", http.StatusBadRequest)
		return
	}
	// database.Products = append(database.Products, newProduct)
	product, err := h.product.Create(repository.TSProducts{
		Title:       req.Title,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
		Description: req.Description,
	}) //database.Store(newProduct)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	utils.SendData(w, product, http.StatusOK)

}
