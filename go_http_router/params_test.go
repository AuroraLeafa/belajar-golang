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

/*
httprouter memiliki kelebihan dibanding ServeMux, yaitu URL yang dinamis,
pada kasus ini dicoba untuk membuat router untuk menampilkan produk dengan Id tertentu di parameter url nya
*/

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		text := "Product " + id
		fmt.Fprint(w, text)
	})
	req := httptest.NewRequest("GET", "http://localhost:8080/products/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	response := rec.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))
}
