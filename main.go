package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/stripe/stripe-go/v79"
)

func main() {
	LoadProducts()
	UpdateProductPrices()

	mux := http.NewServeMux()

	stripe.Key = getSecrets().StripeApi

	mux.Handle("/", http.FileServer(http.Dir("public")))

	mux.HandleFunc("/create-checkout-session", createCheckoutSession)
	mux.HandleFunc("/api/products", getProducts)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(mux)

	addr := "localhost:4242"

	log.Printf("Listening on http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
