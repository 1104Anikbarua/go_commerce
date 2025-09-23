package middlewares

import "ecommerce/config"

type TSMiddleware struct {
	cnf *config.TSConfig
}

func NewMiddleware(config *config.TSConfig) *TSMiddleware {
	return &TSMiddleware{
		cnf: config,
	}
}
