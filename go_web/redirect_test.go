package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func RedirectFrom(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/redirect_to", http.StatusMovedPermanently)
}

func RedirectOut(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://github.com/AuroraLeafa", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect_to", RedirectTo)
	mux.HandleFunc("/redirect_from", RedirectFrom)
	mux.HandleFunc("/redirect_out", RedirectOut)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
