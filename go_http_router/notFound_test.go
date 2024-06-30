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

func TestNotFound(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found!!!")
	})

	req := httptest.NewRequest("GET", "http://localhost:8080/404", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	response := rec.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Not Found!!!", string(body))
}
