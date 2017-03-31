package main

import (
	"log"
	"net/http"

	"encoding/json"

	"io/ioutil"

	"github.com/angelospanag/bazaarvoice-go-sdk/mapping"
	"github.com/angelospanag/bazaarvoice-go-sdk/middleware"
	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/spf13/viper"
)

func main() {

	router := mux.NewRouter()

	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()

	if err != nil {
		log.Println("Config file not found...")
	} else {

		getReviewsForProductChain := alice.New(middleware.ValidateLimit, middleware.ValidateSort)
		router.Handle("/reviews/{productId}", getReviewsForProductChain.ThenFunc(GetReviewsForProduct))

		log.Fatal(http.ListenAndServe(":12345", router))
	}
}

func GetReviewsForProduct(w http.ResponseWriter, r *http.Request) {

	//params := mux.Vars(req)
	//TODO change
	//productId := params["productId"]
	productId := viper.GetString("staging.test_product")
	queryString := viper.GetString("staging.server") + "/data/reviews.json?apiversion=" + viper.GetString("staging.api_version") + "&passkey=" + viper.GetString("staging.api_key") + "&Filter=ProductId:" + productId

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
