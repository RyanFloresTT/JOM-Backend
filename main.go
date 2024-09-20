package main

import (
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v79"
)

func main() {
	stripe.Key = getSecrets().StripeApi

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("POST /create-checkout-session", createCheckoutSession)
	http.HandleFunc("GET /api/products", getProducts)

	addr := "localhost:4242"

	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
