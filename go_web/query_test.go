package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprint(w, "Hello "+name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Reff", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// --------------------Multiple Query Parameters----------------------------
func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	//Disini akan dilakukan test pengiriman beberapa query paramter sekaligus
	//Dengan cara, membuat terus function query dan juga parameter query-nya
	//Jika query parameter tidak diisi, maka akan mengembalikan string kosong
	firstName := r.URL.Query().Get("firstName")
	lastName := r.URL.Query().Get("lastName")

	fmt.Fprintf(w, "%s, %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?firstName=Reff&lastName=Sugg", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// ------------Multiple Value Query Parameter------------
func MultipleParameterValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, strings.Join(names, ","))
}

func TestMultiParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Sugg&name=Reff&name=Cuyy", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
