package middleware

import (
	"net/http"
	"strings"
)

func ValidateSort(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sortParameter := r.URL.Query().Get("sort")

		if len(sortParameter) > 0 {
			sortParameterTokens := strings.Split(sortParameter, ":")

			if len(sortParameterTokens[1]) == 0 || (sortParameterTokens[1] != "asc" && sortParameterTokens[1] != "desc") {
				http.Error(w, http.StatusText(400), http.StatusBadRequest)
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
					http.Error(w, http.StatusText(400), http.StatusBadRequest)
					return
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
