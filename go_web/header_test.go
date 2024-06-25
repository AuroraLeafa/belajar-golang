package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
- Func Header disini digunakan untuk menangkap data-data yang dikirim baik dari client atau server
- Disini menggunakan header yang terdapat pada struct request, lalu method Get untuk mengambil content-type nya
*/

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprintln(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest("POST", "http://localhost:8080", nil)
	request.Header.Set("content-type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// ------------------ Test Response -----------------
func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "ReffKece")
	fmt.Fprint(w, "ok")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	request.Header.Set("content-type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Header.Get("X-Powered-By"))
}
