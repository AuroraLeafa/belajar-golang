package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMapDecode(t *testing.T) {
	jsonRequest := `{"id": "P100", "Name": "Acer", "Price": 2000000}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["Name"])
	fmt.Println(result["Price"])
}

func TestJSONMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P100",
		"Name":  "Acer",
		"Price": 2000000,
	}
	marshal, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}
