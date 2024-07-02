package go_json

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// Membaca File JSON beserta isi datanya
func TestStreamingDecoder(t *testing.T) {
	open, err := os.Open("customer.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(open)

	customer := Customer{}
	decoder.Decode(&customer)
	fmt.Println(customer)
}

// Membuat File JSON beserta isi datanya
func TestStreamingEncoder(t *testing.T) {
	writer, err := os.Create("customer_sample.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Muh",
		MiddleName: "Reff",
		LastName:   "Sugg",
	}

	encoder.Encode(customer)
}
