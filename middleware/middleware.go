package middleware

import (
	"net/http"
)

func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("api-key")
		if header != "sp" {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}
		next.ServeHTTP(w, r)
	})
}
