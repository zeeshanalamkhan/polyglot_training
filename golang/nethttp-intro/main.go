package main

import (
	"fmt"
	"net/http"
)

// handler function to process incoming HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, World!" to the response writer
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Register the handler function for the root URL ("/")
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080 and listen for incoming requests
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// If there's an error starting the server, log it and exit
		fmt.Printf("Error starting server: %v\n", err)
	}
}
