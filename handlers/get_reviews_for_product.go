package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/angelospanag/bazaarvoice-go-sdk/mapping"
	"github.com/google/jsonapi"
	"github.com/spf13/viper"
)

// GetReviewsForProduct Retrieve reviews of a product
func GetReviewsForProduct(w http.ResponseWriter, r *http.Request) {

	//params := mux.Vars(req)
	//TODO change
	//productId := params["productId"]
	productID := viper.GetString("staging.test_product")
	queryString := viper.GetString("staging.server") + "/data/reviews.json?apiversion=" + viper.GetString("staging.api_version") + "&passkey=" + viper.GetString("staging.conversations_api_key") + "&Filter=ProductId:" + productID

	sort := r.URL.Query().Get("sort")
	limit := r.URL.Query().Get("limit")

	// Append 'Sort' parameter if it is requested
	if len(sort) > 0 {
		queryString += "&Sort=" + sort
	}

	// Append 'Limit' parameter if it is requested
	if len(limit) > 0 {
		queryString += "&Limit=" + limit
	}

	responseFromBV, err := http.Get(queryString)

	if err != nil {
		log.Println("Get request error")
	}

	defer responseFromBV.Body.Close()

	reviewsFromBV := mapping.ReviewsFromBV{}

	body, err := ioutil.ReadAll(responseFromBV.Body)

	if err != nil {
		log.Println("Could not read response from BV")
	}

	err = json.Unmarshal(body, &reviewsFromBV)

	if err != nil {
		log.Println("Error unmarshalling reviews from BazaarVoice:", err)
	}

	reviews := mapping.ReviewsMapping(reviewsFromBV)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonapi.MediaType)

	if err := jsonapi.MarshalManyPayload(w, reviews); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
