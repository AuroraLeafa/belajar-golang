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

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})
	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	response := rec.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello World", string(body))
}
