package main

import (
	"fmt"
	_ "fmt"
	"net/http"
	"os"

	"github.com/Man4ct/simple-currency-converter-api/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
	r.POST("/convert", func(c *gin.Context) {
		var requestBody struct {
			Base       string   `json:"base"`
			Amount     int64    `json:"amount"`
			Currencies []string `json:"currencies"`
		}

		// Parse request body
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Call ConvertCurrency function
		result := db.ConvertCurrency(c, requestBody.Base, requestBody.Amount, requestBody.Currencies)

		// Return the result to the client
		c.JSON(http.StatusOK, result)
		// db.ConvertCurrency(c, )
	})
	r.Run()
}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}
