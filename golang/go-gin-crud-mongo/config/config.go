package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client will hold the MongoDB client
var Client *mongo.Client

// ConnectMongoDB initializes the MongoDB client and connects to the database
func ConnectMongoDB() {
	// Create a new MongoDB client and connect to the database
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Adjust the URI as needed
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Set the global client variable
	Client = client

	// Check the connection
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
}

// DisconnectMongoDB closes the connection to the MongoDB server
func DisconnectMongoDB() {
	if err := Client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
}
