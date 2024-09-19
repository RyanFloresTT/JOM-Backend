package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{ID: 1, Name: "Parmesan Crisps", Price: 99.99, Image: "https://res.cloudinary.com/djdtmbpce/image/upload/c_crop,w_250,h_444,ar_9:16/v1726781208/parmesan_crisps_wafcd9.jpg"},
		{ID: 2, Name: "Caesar Crisps", Price: 99.99, Image: "https://res.cloudinary.com/djdtmbpce/image/upload/c_crop,w_250,h_444,ar_9:16/v1726781208/caesar_crisps_gwzhat.jpg"},
		{ID: 3, Name: "Cheese & Garlic Croutons", Price: 99.99, Image: "https://res.cloudinary.com/djdtmbpce/image/upload/c_crop,w_250,h_444,ar_9:16/v1726781208/cheese_garlic_croutons_xqgsa0.jpg"},
		{ID: 4, Name: "Caesar Croutons", Price: 99.99, Image: "https://res.cloudinary.com/djdtmbpce/image/upload/c_crop,w_250,h_444,ar_9:16/v1726781207/casesar_croutons_uuustq.jpg"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/products", getProducts)

	handler := cors.Default().Handler(mux)

	log.Println("Go server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
