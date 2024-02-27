package main

import (
	"net/http"

	"github.com/Man4ct/simple-currency-converter-api/db"
	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}
func convertHandler(c *gin.Context) {
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
}

func latestHandler(c *gin.Context) {
	apiKey, baseURL := loadEnv()
	db.GetLatestCurrency(c, apiKey, baseURL)
}

func getAllCurrencyHandler(c *gin.Context) {
	result := db.GetAllCurrency()
	c.JSON(http.StatusOK, result)
}
