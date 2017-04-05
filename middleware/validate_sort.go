package middleware

import (
	"net/http"
	"strings"

	"github.com/google/jsonapi"
)

// ValidateSort Validate the correct syntax of the 'sort' URL parameter
func ValidateSort(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sortParameter := r.URL.Query().Get("sort")

		if len(sortParameter) > 0 {
			sortParameterTokens := strings.Split(sortParameter, ":")

			if len(sortParameterTokens) == 2 {

				if len(sortParameterTokens[1]) == 0 || (sortParameterTokens[1] != "asc" && sortParameterTokens[1] != "desc") {
					w.Header().Set("Content-Type", jsonapi.MediaType)
					w.WriteHeader(http.StatusBadRequest)
					jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
						Title:  "Validation Error",
						Detail: "Given request body was invalid.",
						Status: "400",
					}})
					return
				}

				if len(sortParameterTokens[0]) > 0 {
					validSortingOptions := []string{"Id", "AuthorId", "CampaignId", "IsRecommended"}
					sortIsValid := false
					for _, validSortingOption := range validSortingOptions {
						if validSortingOption == sortParameterTokens[0] {
							sortIsValid = true
						}
					}

					if !sortIsValid {
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
				}
			} else {
				w.Header().Set("Content-Type", jsonapi.MediaType)
				w.WriteHeader(http.StatusBadRequest)
				jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
					Title:  "Validation Error",
					Detail: "Given request body was invalid.",
					Status: "400",
				}})
				return
			}

		}

		next.ServeHTTP(w, r)
	})
}
