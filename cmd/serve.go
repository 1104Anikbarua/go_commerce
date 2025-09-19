package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	chain := middleware.NewManager()
	initRoute(mux, chain)
	chain.Use(middleware.Logger, middleware.Test)

	fmt.Println("Server Running On Port 8080")
	err := http.ListenAndServe(":8080", middleware.CorsWithPreflight(mux))
	if err != nil {
		fmt.Println("Fail to run the server", err)
	}
	user1.printDetail(100)
	user2.printDetail(50)

}

type TSUser struct {
	Name string
	Age  int
}

var user1 = TSUser{
	Name: "Anik Barua",
	Age:  28,
}
var user2 = TSUser{
	Name: "Tonik Barua",
	Age:  18,
}

func (user TSUser) printDetail(a int) {
	fmt.Println(user.Name)
	fmt.Println(a)
}
