package main_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	server.ListenAndServe()
}
