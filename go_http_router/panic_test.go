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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, err interface{}) {
		fmt.Fprint(writer, "Panic : ", err)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		panic("Ups")
	})
	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	response := rec.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic : Ups", string(body))
}
