package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
XSS adalah suatu masalah keamanan, dimana user kemungkinan akan memasukkan script HTML, CSS atau Javascript
Dan script tersebut jika tidak ditangani maka akan dieksekusi oleh aplikasi,
Developer perlu menangani ini agar keamanan akun user dapat terjaga dari hacker
Untungnya di Golang masalah XSS ini sudah otomatis teratasi berkat template
*/

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "escape.gohtml", map[string]interface{}{
		"Title":  "Auto-Escape",
		"Body":   "<h1>H1 ini akan di-Escape!</h1>",
		"Script": "<script> alert('Anda diHack!'); </script>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//------------------------------------------ AUTO-ESCAPE DISABLED -------------------------------
/*
	Tambahkan fungsi template.HTML(), template.CSS(), atau template.JS()
	pada data yang akan variabel data yang akan diparse
	dengan begitu Golang akan membolehkan script berjalan pada data

	GUNAKAN FUNCTION INI DENGAN HATI-HATI!
*/
func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "escape.gohtml", map[string]interface{}{
		"Title":  "Auto-Escape",
		"Body":   template.HTML("<h1>H1 ini akan di-Escape!</h1>"),
		"Script": template.HTML("<script> alert('Anda diHack!'); </script>"),
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
