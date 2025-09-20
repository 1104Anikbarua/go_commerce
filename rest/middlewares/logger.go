package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Println("This is Logger Middleware")
		log.Println("method->", r.Method, "url->", r.URL.Path, "startTime->", time.Since(startTime))
		next.ServeHTTP(w, r)
	})

}
