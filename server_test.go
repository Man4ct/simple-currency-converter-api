package main_test

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	// Test code will go here

	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
