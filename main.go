package main

import (
	"log"
	"net/http"

	"fmt"

	"encoding/json"

	"github.com/angelospanag/bazaarvoice-go-sdk/domain"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {

	router := mux.NewRouter()

	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Config file not found...")
	} else {
		router.HandleFunc("/reviews/{productId}", GetReviewsForProduct).Methods("GET")
		log.Fatal(http.ListenAndServe(":12345", router))
	}
}

func GetReviewsForProduct(w http.ResponseWriter, req *http.Request) {

	//params := mux.Vars(req)

	//TODO change
	productId := viper.GetString("staging.test_product")
	//productId := params["productId"]

	response, err := http.Get(viper.GetString("staging.server") + "/data/reviews.json?apiversion=" + viper.GetString("staging.api_version") + "&passkey=" + viper.GetString("staging.api_key") + "&Filter=ProductId:" + productId)

	fmt.Println(viper.GetString("staging.server") + "/data/reviews.json?apiversion=" + viper.GetString("staging.api_version") + "&passkey=" + viper.GetString("staging.api_key") + "&Filter=ProductId:" + productId)
	if err != nil {
		log.Fatal("")
	}
	defer response.Body.Close()

	results := domain.Results{}

	if err := json.NewDecoder(response.Body).Decode(&results); err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

}
