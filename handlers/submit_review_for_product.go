package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// ReviewToBePosted A struct holding the data of a review to be posted
type ReviewToBePosted struct {
	ReviewText   string `json:"review_text"`
	Title        string `json:"title"`
	UserNickName string `json:"user_nickname"`
}

// SubmitReviewForProduct Submit a review for a product
func SubmitReviewForProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	decoder := json.NewDecoder(r.Body)
	var review ReviewToBePosted
	err := decoder.Decode(&review)
	if err != nil {
		log.Println()
	}

	productID := params["productID"]

	queryString := viper.GetString("staging.server") + "/data/submitreview.json?apiversion=" + viper.GetString("staging.api_version") + "&passkey=" + viper.GetString("staging.conversations_api_key") + "&Filter=ProductId:" + productID + "&ReviewText=" + review.ReviewText + "&Action=submit&Title=" + review.Title + "&UserNickname=" + review.UserNickName

	resp, err := http.Post(queryString, "application/x-www-form-urlencoded", nil)

	if err != nil {
		log.Println("Cannot post review:" + err.Error())
	}

	defer resp.Body.Close()
}
