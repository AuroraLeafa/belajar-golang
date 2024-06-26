package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Function membuat cookie
func SetCookie(w http.ResponseWriter, r *http.Request) {
	/*
		Cookie.name adalah nama key cookie nya
		cookie.value adalah value cookie, dan dimasukkan pada method Get()
		cookie.path diperlukan untuk menentukan cookie akan digunakan pada URL mana saja
	*/
	cookie := new(http.Cookie)
	cookie.Name = "X-REFF-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success Create Cookie")
}

// Function mengambil/membaca cookie
func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-REFF-Name")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/?name=reff", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, req)
	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s: %s", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/get-cookie", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-REFF-Name"
	cookie.Value = "Reff"
	req.AddCookie(cookie)

	recorder := httptest.NewRecorder()
	GetCookie(recorder, req)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
