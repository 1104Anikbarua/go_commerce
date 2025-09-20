package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
)

func Server() {
	cnf := config.GetConfig()
	rest.Start(cnf)

}
