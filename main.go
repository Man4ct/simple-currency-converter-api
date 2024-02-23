package main

import (
	"fmt"
	_ "fmt"
	"os"
	"time"

	"github.com/Man4ct/simple-currency-converter-api/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// CurrencyResponse represents the structure of the response from the API
type CurrencyResponse struct {
	Success   bool               `json:"success"`
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}

// Currency represents the structure of the currency document in MongoDB
type Currency struct {
	Base      string             `bson:"base"`
	Date      string             `bson:"date"`
	Timestamp int64              `bson:"timestamp"`
	Rates     map[string]float64 `bson:"rates"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		fmt.Println(err)
		panic(err)
	}
	apiKey := os.Getenv("API_KEY")
	baseURL := os.Getenv("BASE_URL")

	// Initialize MongoDB connection
	if err := db.Init(); err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/", handler)
	r.GET("/latest", func(c *gin.Context) {
		db.GetLatestCurrency(c, apiKey, baseURL)
	})
	r.Run()
}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}
