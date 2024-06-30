package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type LogMiddleware struct {
	http.Handler
}

func (middleware LogMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("log middleware - Received Request")
	middleware.Handler.ServeHTTP(w, req)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})
	middleware := LogMiddleware{router}
	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	rec := httptest.NewRecorder()

	middleware.ServeHTTP(rec, req)
	response := rec.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(body))
}
