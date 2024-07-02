package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogMarshal(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestMarshal(t *testing.T) {
	LogMarshal("eko")
	LogMarshal(1)
	LogMarshal(true)
	LogMarshal(nil)
	LogMarshal([]string{"Reff", "Sugg", "Reffs"})
}
