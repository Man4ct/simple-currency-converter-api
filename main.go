package main

import (
	"context"
	"encoding/json"
	"fmt"
	_ "fmt"
	"net/http"
	"os"
	"time"

	"github.com/Man4ct/simple-currency-converter-api/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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
	Base       string             `bson:"base"`
	Date       string             `bson:"date"`
	Timestamp  int64              `bson:"timestamp"`
	Rates      map[string]float64 `bson:"rates"`
	LastUpdate time.Time          `bson:"last_update"`
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
		getLatestCurrency(c, apiKey, baseURL)
	})
	r.Run()
}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}

// TODO: make a function to recalculate the currency value based on the base currency
func getLatestCurrency(c *gin.Context, apiKey, baseURL string) {
	// Construct the URL with the API key as a query parameter
	url := fmt.Sprintf("%s/latest?access_key=%s", baseURL, apiKey)

	// Create a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Send the request
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error making request"})
		return
	}
	defer response.Body.Close()

	// Parse response body into CurrencyResponse struct
	var currencyResponse CurrencyResponse
	err = json.NewDecoder(response.Body).Decode(&currencyResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding response"})
		return
	}
	// Save currency data to MongoDB
	err = saveCurrencyToDB(currencyResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving data to MongoDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": currencyResponse})
}

func saveCurrencyToDB(currencyResponse CurrencyResponse) error {
	client := db.GetClient()
	fmt.Println("Client:", client) // Debugging statement
	// Select database and collection
	database := client.Database("currency")
	collection := database.Collection("currency")

	// Iterate over rates and save each one as a separate document
	for symbol, rate := range currencyResponse.Rates {
		fmt.Println("Symbol:", symbol) // Debugging statement
		fmt.Println("Rate:", rate)     // Debugging statement

		// Create a Currency document
		currency := bson.M{
			"symbol":      symbol,
			"rate":        rate,
			"base":        currencyResponse.Base,
			"date":        currencyResponse.Date,
			"timestamp":   currencyResponse.Timestamp,
			"last_update": time.Now(),
		}

		// Insert the Currency document into the collection
		_, err := collection.InsertOne(context.TODO(), currency)
		if err != nil {
			fmt.Println("Error inserting document:", err) // Debugging statement
			return err
		}
		fmt.Println("Document inserted successfully") // Debugging statement
	}

	return nil
}
