package middleware

import (
	"net/http"
)

type TMiddleware func(handle http.Handler) (handler http.Handler)

type TSManager struct {
	globalMiddlewares []TMiddleware
}

func NewManager() *TSManager {
	mngr := TSManager{
		globalMiddlewares: make([]TMiddleware, 0),
	}
	return &mngr
}
func (mngr *TSManager) Use(middlewares ...TMiddleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *TSManager) With(next http.Handler, middlewares ...TMiddleware) http.Handler {

	n := next

	// for i := len(middlewares) - 1; i >= 0; i-- {
	// 	middleware := middlewares[i]
	// 	n = middleware(n)
	// }
	// return n
	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, globalMiddleware := range mngr.globalMiddlewares {
		n = globalMiddleware(n)
	}
	return n

}
