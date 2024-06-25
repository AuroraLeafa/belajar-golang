package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	/*
		first_name := r.PostFormValue("first_name")
		last_name := r.PostFormValue("last_name")

		Ini adalah cara yang tidak memerlukan parsing secara manual "ParseForm()"
		karena sudah secara otomatis akan diparse oleh method "PostFormValue()"
	*/

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Reff&last_name=Sugg")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//isi dari request.header.add() diatas itu adalah penting harus ada karena sudah jadi ketentuan untuk melakukan form post.
	//tanpa itu form-post tidak akan dieksekusi
	recorder := httptest.NewRecorder()
	FormPost(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))

}
