package go_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, req *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, req *http.Request) {
	file, m, err := req.FormFile("file")
	if err != nil {
		panic(err)
	}
	create, err := os.Create("./uploaded/" + m.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(create, file)
	if err != nil {
		panic(err)
	}
	name := req.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + m.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./uploaded"))))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//----------------------------------- Upload File With Unit Test --------------------
/*
1. Siapkan file yang akan dites menggunakan embed
2. buatlah sebuah writer dengan name yang sesuai pada form.html
3. masukkan file pada writer menggunakan function CreateFormFile() lalu write()
4. Lakukan new Request dengan method post
5. lalu ambil datanya dan print
*/

//go:embed uploaded/Twibbon_DigiClass_2024.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	b := new(bytes.Buffer)
	writer := multipart.NewWriter(b)
	writer.WriteField("name", "Reff")

	file, _ := writer.CreateFormFile("file", "Twibbon_DigiClass_2024.png")
	file.Write(uploadFileTest)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", b)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()

	Upload(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}
