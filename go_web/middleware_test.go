package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (m *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")
	m.Handler.ServeHTTP(w, r)
	fmt.Println("After Execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (e *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Terjadi Error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Terjadi Error: %s", err)
		}
	}()

	e.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprintf(w, "Hello World")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprintf(w, "Hello World")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprintf(w, "Hello World")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Ups")
	})
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
