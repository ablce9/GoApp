package http

import (
	"net/http"
	"log"
)

// LoggingMiddleware ...
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)
	next.ServeHTTP(w, r)
    })
}
