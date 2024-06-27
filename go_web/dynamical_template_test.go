package go_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates1 embed.FS

//--------------------------- USING MAP INTERFACE -----------------------------------------

func TemplateDataInterface(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates1, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map Interface",
		"Name":  "Reff",
		"Address": map[string]interface{}{
			"Street": "Jalan Bogor",
		},
	})
}

func TestTemplateDataInterface(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataInterface(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// --------------------------- USING STRUCT -----------------------------------------
type Address struct {
	Street string
	City   string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates1, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Reff Sugg",
		Address: Address{
			Street: "Jalan Belum Ada",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
