package middleware

import (
	"net/http"
	"log"
)

type Response struct {
	statusCode int
	value interface{}
}

func JsonResponseMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		log.Fatal(r.Context().Value("response"))
	})

}
