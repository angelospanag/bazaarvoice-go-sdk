package middleware

import (
	"net/http"
	"regexp"

	"github.com/google/jsonapi"
)

// ValidateLimit Validate that the 'limit' URL parameter is a number
func ValidateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		limit := r.URL.Query().Get("limit")

		if len(limit) > 0 {
			match, _ := regexp.MatchString("^[0-9]+$", limit)

			if !match {

				// An error has come up in your code, so set an appropriate status, and serialize the error.
				w.Header().Set("Content-Type", jsonapi.MediaType)
				w.WriteHeader(http.StatusBadRequest)
				jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
					Title:  "Validation Error",
					Detail: "Given request body was invalid.",
					Status: "400",
				}})
				return

				//http.Error(w, http.StatusText(400), http.StatusBadRequest)
				//return
			}
		}
		next.ServeHTTP(w, r)
	})
}
