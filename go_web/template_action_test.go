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
var templates_action embed.FS

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates_action, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "",
		Address: Address{
			Street: "Jalan Belum Ada",
		},
	})
}

func TestTemplateAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//----------------------------  COMPARATOR --------------------

/*
Operator Perbandingan:
	- eq (equal) -> arg1 == arg2
	- ne (Not Equal) -> arg1 != arg2
	- lt (Less Than) -> arg1 < arg2
	- le (Less Equual) -> arg1 <= arg2
	- gt (greater than) -> arg1 > arg2
	- ge (greater equal) -> arg1 >= arg2
*/

func TemplateActionComparator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates_action, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title":      "Comparator Test",
		"FinalValue": 90,
	})
}

func TestTemplateComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionComparator(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// ---------------------------- Looping (Using Range) --------------------
func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates_action, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Range Test",
		"Hobbies": []string{
			"Gaming", "Reading", "Music", "Coding", "Movies",
		},
	})
}

func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// ---------------------------- Looping (Using Range) --------------------
func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates_action, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Title": "Address Test",
		"Address": Address{
			"Puncak", "Bogor",
		},
	})
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
