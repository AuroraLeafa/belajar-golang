package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		MiddleName: "Smith",
		LastName:   "Doe",
		Hobbies:    []string{"Gaming", "Reading", "Movies", "Coding"},
	}

	marshal, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"John","MiddleName":"Smith","LastName":"Doe","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Movies","Coding"]}`
	jsonBytes := []byte(jsonString)
	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

func TestJSONArrayComplexEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		MiddleName: "Smith",
		LastName:   "Doe",
		Hobbies:    []string{"Gaming", "Reading", "Movies", "Coding"},
		Addresses: []Address{
			{
				Street:     "Jalan Tol",
				Country:    "Indonesia",
				PostalCode: 124214,
			},
			{
				Street:     "Jalan Jalan",
				Country:    "Wakanda",
				PostalCode: 124999,
			},
		},
	}

	marshal, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"John","MiddleName":"Smith","LastName":"Doe","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Movies","Coding"],"Addresses":[{"Street":"Jalan Tol","Country":"Indonesia","PostalCode":124214},{"Street":"Jalan Jalan","Country":"Wakanda","PostalCode":124999}]}`
	jsonBytes := []byte(jsonString)
	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Tol","Country":"Indonesia","PostalCode":124214},{"Street":"Jalan Jalan","Country":"Wakanda","PostalCode":124999}]`
	jsonBytes := []byte(jsonString)
	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}
