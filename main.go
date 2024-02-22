package main

import (
	"fmt"
	"net/http"

	_ "github.com/gin-gonic/gin"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}