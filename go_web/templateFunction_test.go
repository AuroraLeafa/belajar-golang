package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ". My Name is, " + myPage.Name
}

// Menggunakan Function saat parsing
func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("Functions").Parse(`{{.SayHello "Cuyy"}}`))
	t.ExecuteTemplate(writer, "Functions", MyPage{
		Name: "Reff",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// ----------------------------------- Global Function --------------------------
/*
	- registrasi function menggunakan function t.Funcs()
	- Setiap Function dibuat sebagai anonymous func dalam map
	- Setelah Function dibuat baru function dapat di parse ke template
*/
func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FunctionGlobal")
	t = t.Funcs(map[string]interface{}{
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
	})
	t = template.Must(t.Parse("{{upper .Name}} {{lower .Name}}"))
	t.ExecuteTemplate(writer, "FunctionGlobal", MyPage{
		Name: "Reff",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// ----------------------------------- Pipeline Function --------------------------
/*
	Pipeline adalah function untuk mengirimkan hasil dari function A ke function lainnya
	Pada kasus ini dilakukan test pipeline yaitu
		- Function Upper akan menerima parameter "Reff"
		- Lalu dilakukan Pipeline (Simbol | ) sehingga hasil dari function upper
  			akan dikirim ke function setelahnya
		- Function SayHello kemudian menerima hasil return dari function Upper
			sehingga function SayHello akan memberikan Name yaitu "REFF"
*/
func TemplateFunctionPipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FunctionGlobal")
	t = t.Funcs(map[string]interface{}{
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
		"SayHello": func(s string) string {
			return "Hello, My Name is, " + s
		},
	})
	t = template.Must(t.Parse("{{upper .Name | SayHello}}"))
	t.ExecuteTemplate(writer, "FunctionGlobal", MyPage{
		Name: "Reff",
	})
}

func TestTemplateFunctionPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipeline(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
