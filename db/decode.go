package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func decodeMany(documents *[]bson.M, result *mongo.Cursor) {
	// Iterate over the cursor to access each document
	for result.Next(context.TODO()) {
		var document bson.M
		if err := result.Decode(&document); err != nil {
			fmt.Println("Error decoding document:", err)
			continue
		}
		*documents = append(*documents, document)
	}
}

func decodeSingle(document *bson.M, result *mongo.SingleResult) {
	if err := result.Decode(document); err != nil {
		fmt.Println("Error decoding document:", err)
	}
}
