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

func TestRouterNamedParameter(t *testing.T) {
	/*
		Dengan named parameter dapat membuat 1 atau lebih parameter pada URL
		dengan cara menambahkan terus "/:NamaParams" sebanyak yang diperlukan
	*/
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId",
		func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
			text := "Product " + params.ByName("id") + " Item " + params.ByName("itemId")
			fmt.Fprint(w, text)
		})
	req := httptest.NewRequest("GET", "http://localhost:8080/products/1/items/2", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	response := rec.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 2", string(body))
}

func TestRouterCatchAll(t *testing.T) {
	/*
		Catch all param dengan cara "/*NamaParams" membuat golang dapat
		menangkap semua parameter setelah url yang telah ditentukan
	*/
	router := httprouter.New()
	router.GET("/images/*image",
		func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
			text := "Image : " + params.ByName("image")
			fmt.Fprint(w, text)
		})
	req := httptest.NewRequest("GET", "http://localhost:8080/images/small/profile.png", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	response := rec.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image : /small/profile.png", string(body))
}
