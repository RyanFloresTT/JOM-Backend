package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type secrets struct {
	StripeApi string `json:"stripeApi"`
}

func getSecrets() secrets {
	var s secrets
	file, err := os.Open("secrets.json")
	if err != nil {
		fmt.Println("Error opening secrets file:", err)
		return s
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&s)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	return s
}
