package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is required")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}
func TestResponseCodeInvalid(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, req)
	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
	fmt.Println("Status Code :", result.StatusCode)
	fmt.Println("Current Status:", result.Status)
}

func TestResponseCodeValid(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/?name=Reff", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, req)
	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
	fmt.Println("Status Code :", result.StatusCode)
	fmt.Println("Current Status:", result.Status)
}
