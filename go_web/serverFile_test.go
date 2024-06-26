package go_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

/*
Method ServeFile digunakan untuk menampilkan hanya file/halaman tertentu
tanpa perlu membaca semua file/halaman pada folder tersebut

Pada kasus ini aplikasi akan menampilkan hanya halaman Ok atau notfound saat kondisi terpenuhi
dan file/halaman lainnya yang berada di direktori tersebut tidak bisa diakses
*/
func ServeFile(w http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Get("name") != "" {
		http.ServeFile(w, req, "./resources/ok.html")
	} else {
		http.ServeFile(w, req, "./resources/notFound.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// ------------------------ DENGAN GO-EMBED -------------------------------

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notFound.html
var resourceNotFound string

func ServeFileEmbed(w http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourceOk)
	} else {
		fmt.Fprint(w, resourceNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
