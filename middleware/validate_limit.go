package middleware

import (
	"net/http"
	"regexp"
)

// ValidateLimit Validate that the 'limit' URL parameter is a number
func ValidateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		limit := r.URL.Query().Get("limit")

		if len(limit) > 0 {
			match, _ := regexp.MatchString("^[0-9]+$", limit)

			if !match {
				http.Error(w, http.StatusText(400), http.StatusBadRequest)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
