package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Server() {
	chain := middleware.NewManager()
	chain.Use(middleware.Preflight, middleware.Cors, middleware.Test, middleware.Logger)
	mux := http.NewServeMux()
	wrapedMux := chain.WrapMux(mux)

	initRoute(mux, chain)
	fmt.Println("Server Running On Port 8080")
	err := http.ListenAndServe(":8080", wrapedMux)
	if err != nil {
		fmt.Println("Fail to run the server", err)
	}

}
