package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/stripe/stripe-go/v79"
)

func main() {
	mux := http.NewServeMux()

	stripe.Key = getSecrets().StripeApi

	mux.Handle("/", http.FileServer(http.Dir("public")))

	mux.HandleFunc("/create-checkout-session", createCheckoutSession)
	mux.HandleFunc("/api/products", getProducts)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	}).Handler(mux)

	addr := "localhost:4242"

	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
