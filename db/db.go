package db

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Init initializes the MongoDB connection
func Init() error {
	ctx := context.TODO()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB")

	// Create the collection if it does not exist
	if err := createCollection(); err != nil {
		return err
	}

	return nil
}

// GetClient returns the MongoDB client
func GetClient() *mongo.Client {
	return client
}

func createCollection() error {
	// Select the database
	database := client.Database("currency")
	fmt.Println("Database:", database) // Debugging statement
	fmt.Println(database)              // Debugging statement
	// Create the collection options
	options := options.CreateCollection()
	options.SetValidator(bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"symbol", "rate"},
			"properties": bson.M{
				"symbol": bson.M{
					"bsonType":    "string",
					"description": "must be a string and is required",
					"minLength":   1,
				},
				"rate": bson.M{
					"bsonType":    "double",
					"description": "must be a double and is required",
				},
			},
		},
	})

	// Create the collection
	err := database.CreateCollection(context.TODO(), "currency", options)
	if err != nil {
		fmt.Println("Error creating collection:", err) // Debugging statement
		// Check if collection already exists
		if !strings.Contains(err.Error(), "NamespaceExists") {
			return err
		}
	}

	return nil
}
