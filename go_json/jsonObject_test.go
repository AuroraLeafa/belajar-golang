package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode int
}

// --------------------------------------------- ENCODE JSON -------------------------
// Marshal -> Mengubah struct menjadi JSON (ENCODE)
func TestEncodeJSON(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		MiddleName: "Doe",
		LastName:   "Smith",
		Age:        34,
		Married:    false,
	}
	marshal, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}

// --------------------------------------------- DECODE JSON -------------------------
func TestDecodeJSON(t *testing.T) {
	//Unmarshal() -> Mengubah JSON menjadi Struct
	jsonString := `{"FirstName" : "Reff", "MiddleName" : "Sugg", "LastName" : "Riffs", "Age": 34, "Married": true}`
	jsonBytes := []byte(jsonString)
	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}
