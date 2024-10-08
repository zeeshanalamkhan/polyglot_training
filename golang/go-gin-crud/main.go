// Package main is the entry point for the application.
package main

import (
	"github.com/gitish/polyglot_training/routes"
)

// main initializes and starts the web server.
func main() {
	router := routes.SetupRouter()

	// Start the Gin server on port 8080
	router.Run(":8080")
}
