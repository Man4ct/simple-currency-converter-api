package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongoDB initializes the MongoDB connection
func InitMongoDB() (*mongo.Client, error) {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB")
	return client, nil
}
