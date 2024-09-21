package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/price"
)

type Product struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Image   string  `json:"image"`
	PriceID string  `json:"priceID"`
	Price   float64 `json:"price"`
}

var products []Product

func LoadProducts() {
	file, err := os.Open("products.json")
	if err != nil {
		log.Fatalf("Failed to open products file: %v", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&products); err != nil {
		log.Fatalf("Failed to decode products: %v", err)
	}
}

func SaveProducts() {
	file, err := os.Create("products.json")
	if err != nil {
		log.Fatalf("Failed to create products file: %v", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(products); err != nil {
		log.Fatalf("Failed to encode products: %v", err)
	}
}

func UpdateProductPrices() {
	stripe.Key = getSecrets().StripeApi
	for i, product := range products {
		p, err := price.Get(product.PriceID, nil)
		if err != nil {
			log.Printf("Error retrieving price for Product ID %d: %v", product.ID, err)
			continue
		}

		newPrice := p.UnitAmountDecimal / 100.0
		if newPrice != product.Price {
			products[i].Price = newPrice
		}
	}

	SaveProducts()
}
