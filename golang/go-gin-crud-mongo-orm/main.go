// Package main is the entry point for the application.
package main

import (
	"github.com/gitish/polyglot_training/config"
	"github.com/gitish/polyglot_training/routes"
)

// main initializes and starts the web server.
func main() {
	// Connect to MongoDB
	config.ConnectMongoDB()

	// Ensure the connection is closed when the application exits
	defer config.DisconnectMongoDB()

	// Set up routes and start the server
	router := routes.SetupRouter(config.Client) // Pass the MongoDB client
	router.Run(":8080")                         // Start the server on port 8080
}
