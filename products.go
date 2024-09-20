package main

type Product struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	PriceID string `json:"priceID"`
}

var products = []Product{
	{
		ID:      1,
		Name:    "Parmesan Crisps",
		Image:   "https://res.cloudinary.com/djdtmbpce/image/upload/c_crop,w_250,h_444,ar_9:16/v1726781208/parmesan_crisps_wafcd9.jpg",
		PriceID: "price_1Q0z10P4AZbXqL5ZcgdDFZ4u",
	},
	{
		ID:      2,
		Name:    "Caesar Crisps",
		Image:   "https://res.cloudinary.com/djdtmbpce/image/upload/c_crop,w_250,h_444,ar_9:16/v1726781208/caesar_crisps_gwzhat.jpg",
		PriceID: "price_1Q0yvbP4AZbXqL5ZDQnXErlC",
	},
	{
		ID:      3,
		Name:    "Cheese & Garlic Croutons",
		Image:   "https://res.cloudinary.com/djdtmbpce/image/upload/c_crop,w_250,h_444,ar_9:16/v1726781208/cheese_garlic_croutons_xqgsa0.jpg",
		PriceID: "price_1Q0torP4AZbXqL5ZqKMsKdSI",
	},
	{
		ID:      4,
		Name:    "Caesar Croutons",
		Image:   "https://res.cloudinary.com/djdtmbpce/image/upload/c_crop,w_250,h_444,ar_9:16/v1726781207/casesar_croutons_uuustq.jpg",
		PriceID: "price_1Q1D5qP4AZbXqL5ZB6uv9oSm",
	},
}
