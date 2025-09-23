package review

import (
	"ecommerce/utils"
	"net/http"
)

func (h *TSNewHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, "Reviews is coming", http.StatusOK)
}
