package main

import (
	_ "context"
	"fmt"
	"net/http"

	"github.com/Man4ct/simple-currency-converter-api/db"
	_ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db.InitMongoDB()
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
