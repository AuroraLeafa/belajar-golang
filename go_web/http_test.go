package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	//Unit Test Server tanpa perlu menjalankan server ataupun cek pada web
	request, err := http.NewRequest("GET", "http://127.0.0.1:8080/hello", nil)
	if err != nil {
		panic(err)
	}
	recorder := httptest.NewRecorder()
	HelloHandler(recorder, request)
	recorder.Result()
	body, _ := io.ReadAll(recorder.Body)
	fmt.Println(string(body))
}
