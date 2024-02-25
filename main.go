package main

import (
	"fmt"
	_ "fmt"
	"os"

	"github.com/Man4ct/simple-currency-converter-api/db"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize MongoDB connection
	if err := db.Init(); err != nil {
		panic(err)
	}
	startServer()
}

func loadEnv() (string, string) {
	if err := godotenv.Load("./.env"); err != nil {
		fmt.Println(err)
		panic(err)
	}

	apiKey := os.Getenv("API_KEY")
	baseURL := os.Getenv("BASE_URL")

	return apiKey, baseURL
}
