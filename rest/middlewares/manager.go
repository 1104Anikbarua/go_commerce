package middleware

import (
	"fmt"
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

	h := next
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h

}

func (mngr *TSManager) WrapMux(handler http.Handler) http.Handler {
	h := handler
	fmt.Println("GLOBAL_MIDDLEWARE", mngr.globalMiddlewares)
	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h)
	}
	return h
}
