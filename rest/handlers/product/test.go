package product

import (
	"log"
	"net/http"
)

func (h *TSNewHandler) Test(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT")
}
