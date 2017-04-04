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

		getReviewsForProductChain := alice.New(middleware.ValidateLimit, middleware.ValidateSort)
		router.Handle("/product/{productId}/reviews", getReviewsForProductChain.ThenFunc(handlers.GetReviewsForProduct))

		log.Fatal(http.ListenAndServe(":12345", router))
	}
}
