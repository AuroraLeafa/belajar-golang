package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// JSON Tag ini diperlukan jika json yang diterima tidak sesuai dengan ketentuan Golang (Pascalcase)
type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTagEncode(t *testing.T) {
	product := Product{
		Id:       "P-0001",
		Name:     "Acer",
		ImageURL: "example.com/image.png",
	}

	marshal, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P-0001","name":"Acer","Image_URL":"example.com/image.png"}`
	jsonBytes := []byte(jsonString)
	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
