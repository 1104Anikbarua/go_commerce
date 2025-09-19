package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Server() {
	cnf := config.GetConfig()
	chain := middleware.NewManager()
	chain.Use(middleware.Preflight, middleware.Cors, middleware.Test, middleware.Logger)
	mux := http.NewServeMux()
	wrapedMux := chain.WrapMux(mux)

	initRoute(mux, chain)
	port := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server Running On Port", port)
	err := http.ListenAndServe(port, wrapedMux)
	if err != nil {
		fmt.Println("Fail to run the server", err)
		os.Exit(1)
	}

}
