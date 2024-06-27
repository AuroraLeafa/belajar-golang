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

func SimpleHTML(w http.ResponseWriter, req *http.Request) {
	templateText := `<!DOCTYPE html><body>{{.}}</body></html>`
	t := template.Must(template.New("SIMPLE").Parse(templateText))

	err := t.ExecuteTemplate(w, "SIMPLE", "Hello World")
	if err != nil {
		panic(err)
	}
}

func TestTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// -------------------------- Template Using File -------------------------
func SimpleHTMLFile(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	err := t.ExecuteTemplate(w, "simple.gohtml", "Hasil Parsing")
	if err != nil {
		fmt.Println("Parse Failed!")
	}
}

func TestParsingFile(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	SimpleHTMLFile(rec, req)
	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

// ---------------------- Template Using Whole Directory ----------------------------
func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "simple.gohtml", "INI PARSING dengan Glob")
}

func TestTemplateDirectory(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateDirectory(rec, req)
	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

// --------------------- Parsing with Embed --------------------

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	t.ExecuteTemplate(w, "simple.gohtml", "INI PARSING Dengan Embed")
}

func TestTemplateWithEmbed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateEmbed(rec, req)
	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}
