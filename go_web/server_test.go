package go_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi, My Name Is Reff~")
	})
	mux.HandleFunc("/hi/indo/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Halo, nama saya Reff~")
	})
	mux.HandleFunc("/yuhu", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "YUUUHUUUUUUU")
	})
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
