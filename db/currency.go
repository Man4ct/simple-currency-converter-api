package db

import (
	"context"
	"encoding/json"
	"fmt"
	_ "fmt"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

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

func GetLatestCurrency(c *gin.Context, apiKey, baseURL string) {
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

	// // Save currency data to MongoDB if the collection is empty
	err = saveCurrencyToDB(currencyResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving data to MongoDB"})
		return
	}

	// Update currency data to MongoDB
	err = updateCurrenciesInDB(currencyResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving data to MongoDB"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": currencyResponse})
}

func saveCurrencyToDB(currencyResponse CurrencyResponse) error {
	currencies := map[string]string{
		"AED": "United Arab Emirates Dirham",
		"AFN": "Afghan Afghani",
		"ALL": "Albanian Lek",
		"AMD": "Armenian Dram",
		"ANG": "Netherlands Antillean Guilder",
		"AOA": "Angolan Kwanza",
		"ARS": "Argentine Peso",
		"AUD": "Australian Dollar",
		"AWG": "Aruban Florin",
		"AZN": "Azerbaijani Manat",
		"BAM": "Bosnia-Herzegovina Convertible Mark",
		"BBD": "Barbadian Dollar",
		"BDT": "Bangladeshi Taka",
		"BGN": "Bulgarian Lev",
		"BHD": "Bahraini Dinar",
		"BIF": "Burundian Franc",
		"BMD": "Bermudan Dollar",
		"BND": "Brunei Dollar",
		"BOB": "Bolivian Boliviano",
		"BRL": "Brazilian Real",
		"BSD": "Bahamian Dollar",
		"BTC": "Bitcoin",
		"BTN": "Bhutanese Ngultrum",
		"BWP": "Botswanan Pula",
		"BYN": "New Belarusian Ruble",
		"BYR": "Belarusian Ruble",
		"BZD": "Belize Dollar",
		"CAD": "Canadian Dollar",
		"CDF": "Congolese Franc",
		"CHF": "Swiss Franc",
		"CLF": "Chilean Unit of Account (UF)",
		"CLP": "Chilean Peso",
		"CNY": "Chinese Yuan",
		"COP": "Colombian Peso",
		"CRC": "Costa Rican Colón",
		"CUC": "Cuban Convertible Peso",
		"CUP": "Cuban Peso",
		"CVE": "Cape Verdean Escudo",
		"CZK": "Czech Republic Koruna",
		"DJF": "Djiboutian Franc",
		"DKK": "Danish Krone",
		"DOP": "Dominican Peso",
		"DZD": "Algerian Dinar",
		"EGP": "Egyptian Pound",
		"ERN": "Eritrean Nakfa",
		"ETB": "Ethiopian Birr",
		"EUR": "Euro",
		"FJD": "Fijian Dollar",
		"FKP": "Falkland Islands Pound",
		"GBP": "British Pound Sterling",
		"GEL": "Georgian Lari",
		"GGP": "Guernsey Pound",
		"GHS": "Ghanaian Cedi",
		"GIP": "Gibraltar Pound",
		"GMD": "Gambian Dalasi",
		"GNF": "Guinean Franc",
		"GTQ": "Guatemalan Quetzal",
		"GYD": "Guyanaese Dollar",
		"HKD": "Hong Kong Dollar",
		"HNL": "Honduran Lempira",
		"HRK": "Croatian Kuna",
		"HTG": "Haitian Gourde",
		"HUF": "Hungarian Forint",
		"IDR": "Indonesian Rupiah",
		"ILS": "Israeli New Sheqel",
		"IMP": "Manx pound",
		"INR": "Indian Rupee",
		"IQD": "Iraqi Dinar",
		"IRR": "Iranian Rial",
		"ISK": "Icelandic Króna",
		"JEP": "Jersey Pound",
		"JMD": "Jamaican Dollar",
		"JOD": "Jordanian Dinar",
		"JPY": "Japanese Yen",
		"KES": "Kenyan Shilling",
		"KGS": "Kyrgystani Som",
		"KHR": "Cambodian Riel",
		"KMF": "Comorian Franc",
		"KPW": "North Korean Won",
		"KRW": "South Korean Won",
		"KWD": "Kuwaiti Dinar",
		"KYD": "Cayman Islands Dollar",
		"KZT": "Kazakhstani Tenge",
		"LAK": "Laotian Kip",
		"LBP": "Lebanese Pound",
		"LKR": "Sri Lankan Rupee",
		"LRD": "Liberian Dollar",
		"LSL": "Lesotho Loti",
		"LTL": "Lithuanian Litas",
		"LVL": "Latvian Lats",
		"LYD": "Libyan Dinar",
		"MAD": "Moroccan Dirham",
		"MDL": "Moldovan Leu",
		"MGA": "Malagasy Ariary",
		"MKD": "Macedonian Denar",
		"MMK": "Myanma Kyat",
		"MNT": "Mongolian Tugrik",
		"MOP": "Macanese Pataca",
		"MRU": "Mauritanian Ouguiya",
		"MUR": "Mauritian Rupee",
		"MVR": "Maldivian Rufiyaa",
		"MWK": "Malawian Kwacha",
		"MXN": "Mexican Peso",
		"MYR": "Malaysian Ringgit",
		"MZN": "Mozambican Metical",
		"NAD": "Namibian Dollar",
		"NGN": "Nigerian Naira",
		"NIO": "Nicaraguan Córdoba",
		"NOK": "Norwegian Krone",
		"NPR": "Nepalese Rupee",
		"NZD": "New Zealand Dollar",
		"OMR": "Omani Rial",
		"PAB": "Panamanian Balboa",
		"PEN": "Peruvian Nuevo Sol",
		"PGK": "Papua New Guinean Kina",
		"PHP": "Philippine Peso",
		"PKR": "Pakistani Rupee",
		"PLN": "Polish Zloty",
		"PYG": "Paraguayan Guarani",
		"QAR": "Qatari Rial",
		"RON": "Romanian Leu",
		"RSD": "Serbian Dinar",
		"RUB": "Russian Ruble",
		"RWF": "Rwandan Franc",
		"SAR": "Saudi Riyal",
		"SBD": "Solomon Islands Dollar",
		"SCR": "Seychellois Rupee",
		"SDG": "South Sudanese Pound",
		"SEK": "Swedish Krona",
		"SGD": "Singapore Dollar",
		"SHP": "Saint Helena Pound",
		"SLE": "Sierra Leonean Leone",
		"SLL": "Sierra Leonean Leone",
		"SOS": "Somali Shilling",
		"SRD": "Surinamese Dollar",
		"STD": "São Tomé and Príncipe Dobra",
		"SVC": "Salvadoran Colón",
		"SYP": "Syrian Pound",
		"SZL": "Swazi Lilangeni",
		"THB": "Thai Baht",
		"TJS": "Tajikistani Somoni",
		"TMT": "Turkmenistani Manat",
		"TND": "Tunisian Dinar",
		"TOP": "Tongan Paʻanga",
		"TRY": "Turkish Lira",
		"TTD": "Trinidad and Tobago Dollar",
		"TWD": "New Taiwan Dollar",
		"TZS": "Tanzanian Shilling",
		"UAH": "Ukrainian Hryvnia",
		"UGX": "Ugandan Shilling",
		"USD": "United States Dollar",
		"UYU": "Uruguayan Peso",
		"UZS": "Uzbekistan Som",
		"VEF": "Venezuelan Bolívar Fuerte",
		"VES": "Sovereign Bolivar",
		"VND": "Vietnamese Dong",
		"VUV": "Vanuatu Vatu",
		"WST": "Samoan Tala",
		"XAF": "CFA Franc BEAC",
		"XAG": "Silver (troy ounce)",
		"XAU": "Gold (troy ounce)",
		"XCD": "East Caribbean Dollar",
		"XDR": "Special Drawing Rights",
		"XOF": "CFA Franc BCEAO",
		"XPF": "CFP Franc",
		"YER": "Yemeni Rial",
		"ZAR": "South African Rand",
		"ZMK": "Zambian Kwacha (pre-2013)",
		"ZMW": "Zambian Kwacha",
		"ZWL": "Zimbabwean Dollar",
	}
	client := GetClient()
	fmt.Println("Client:", client) // Debugging statement
	// Select database and collection
	database := client.Database("currency")
	collection := database.Collection("currency")

	// Check if the collection is empty
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		// Collection is not empty, do not save data
		return nil
	}

	// Iterate over rates and save each one as a separate document
	for symbol, rate := range currencyResponse.Rates {
		fmt.Println("Symbol:", symbol) // Debugging statement
		fmt.Println("Rate:", rate)     // Debugging statement
		name := "Test Name"
		currency, ok := currencies[symbol]
		if ok {
			name = currency
		}
		// Create a Currency document
		currencyDocument := bson.M{
			"symbol":    symbol,
			"rate":      rate,
			"base":      currencyResponse.Base,
			"date":      currencyResponse.Date,
			"timestamp": currencyResponse.Timestamp,
			"name":      name,
			// "name":
		}

		// Insert the Currency document into the collection
		_, err := collection.InsertOne(context.TODO(), currencyDocument)
		if err != nil {
			fmt.Println("Error inserting document:", err) // Debugging statement
			return err
		}
		fmt.Println("Document inserted successfully") // Debugging statement
	}

	return nil
}

func updateCurrenciesInDB(currencyResponse CurrencyResponse) error {
	client := GetClient()
	database := client.Database("currency")
	collection := database.Collection("currency")

	for symbol, rate := range currencyResponse.Rates {
		filter := bson.M{"symbol": symbol}
		update := bson.M{"$set": bson.M{"rate": rate, "updated_at": time.Now()}}

		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			fmt.Println("Error updating document:", err) // Debugging statement
			return err
		}

		fmt.Println("Document updated successfully") // Debugging statement
	}

	return nil
}

func ConvertCurrency(c *gin.Context, base string, amount int64, curencies []string) gin.H {
	client := GetClient()
	database := client.Database("currency")
	collection := database.Collection("currency")

	baseCurrency := collection.FindOne(context.TODO(), bson.M{"symbol": base})
	result, err := collection.Find(context.TODO(), bson.M{"symbol": bson.M{"$in": curencies}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding currency"})
		// return err
	}
	// Access the single document from baseCurrency
	var baseDocument bson.M
	decodeSingle(&baseDocument, baseCurrency)

	var documents []bson.M
	decodeMany(&documents, result)

	// Check if there were any errors during iteration
	if err := result.Err(); err != nil {
		fmt.Println("Error during iteration:", err)
	}

	baseRate := baseDocument["rate"].(float64)

	for _, document := range documents {
		rate := document["rate"].(float64)
		convertedAmount := (1 / baseRate) * rate * float64(amount)

		// Round the converted amount to four digits after the decimal point
		roundedAmount := math.Round(convertedAmount*10000) / 10000

		// Update the document with the rounded converted amount
		document["converted_amount"] = roundedAmount

	}
	return gin.H{
		"base":  baseDocument,
		"rates": documents,
	}
}

func GetAllCurrency() []bson.M {
	client := GetClient()
	database := client.Database("currency")
	collection := database.Collection("currency")

	// Find all documents in the collection
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("Error finding documents:", err)
		return nil
	}

	var curencies []bson.M
	decodeMany(&curencies, cursor)
	return curencies
}
