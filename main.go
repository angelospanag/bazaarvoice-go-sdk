package main

import (
	"log"
	"net/http"

	"github.com/angelospanag/bazaarvoice-go-sdk/handlers"
	"github.com/angelospanag/bazaarvoice-go-sdk/middleware"
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

		getReviewsForProductChain := alice.New(middleware.ValidateProductID, middleware.ValidateLimit, middleware.ValidateSort)
		submitReviewForProductChain := alice.New(middleware.ValidateProductID)

		router.Handle("/product/{productID}/reviews", getReviewsForProductChain.ThenFunc(handlers.GetReviewsForProduct)).Methods("GET")
		router.Handle("/product/{productID}/review", submitReviewForProductChain.ThenFunc(handlers.SubmitReviewForProduct)).Methods("POST")

		log.Fatal(http.ListenAndServe(":12345", router))
	}
}
