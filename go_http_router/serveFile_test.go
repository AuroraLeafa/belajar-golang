package main

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	sub, err := fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}
	router.ServeFiles("/files/*filepath", http.FS(sub))
	req := httptest.NewRequest("GET", "http://localhost:8080/files/hello.txt", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	response := rec.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello World", string(body))
}
