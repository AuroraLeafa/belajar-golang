package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Please provide a file name")
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(w, r, "./uploaded/"+file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
