package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	UpdateProductPrices()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func createCheckoutSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var checkoutItems checkoutRequest
	err := json.NewDecoder(r.Body).Decode(&checkoutItems)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var lineItems []*stripe.CheckoutSessionLineItemParams
	var sessionMode stripe.CheckoutSessionMode
	sessionMode = stripe.CheckoutSessionModePayment

	for _, item := range checkoutItems.Items {
		lineItem := &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.PriceID),
			Quantity: stripe.Int64(int64(item.Quantity)),
		}

		if item.Subscription {
			sessionMode = stripe.CheckoutSessionModeSubscription
		}

		lineItems = append(lineItems, lineItem)
	}

	params := &stripe.CheckoutSessionParams{
		LineItems:  lineItems,
		Mode:       stripe.String(string(sessionMode)),
		SuccessURL: stripe.String("http://localhost:3000/cart?success=true"),
		CancelURL:  stripe.String("http://localhost:3000/cart?canceled=true"),
	}

	s, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
		http.Error(w, "Failed to create checkout session", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"url": s.URL})
}

type checkoutItem struct {
	PriceID      string `json:"priceID"`
	Quantity     int    `json:"quantity"`
	Subscription bool   `json:"subscription"`
}

type checkoutRequest struct {
	Items []checkoutItem `json:"items"`
}
