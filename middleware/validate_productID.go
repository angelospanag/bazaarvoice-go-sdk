package middleware

import (
	"net/http"
	"regexp"

	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
)

// ValidateProductID Validate if a product ID has a correct format
func ValidateProductID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		productID := params["productID"]

		matched, _ := regexp.MatchString("^[0-9]+$", productID)
		if !matched {
			// An error has come up in your code, so set an appropriate status, and serialize the error.
			w.Header().Set("Content-Type", jsonapi.MediaType)
			w.WriteHeader(http.StatusBadRequest)
			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "Validation Error",
				Detail: "Given request body was invalid.",
				Status: "400",
			}})
			return
		}

		next.ServeHTTP(w, r)
	})
}
