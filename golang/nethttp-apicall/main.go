package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Todo struct represents the structure of a TODO item from the API response
type Todo struct {
	UserID    int    `json:"userId"`    // ID of the user who created the todo
	ID        int    `json:"id"`        // ID of the todo item
	Title     string `json:"title"`     // Title of the todo item
	Completed bool   `json:"completed"` // Status indicating whether the todo is completed
}

// Define the API endpoint URL
const URL = "https://jsonplaceholder.typicode.com/todos/1"

// main is the entry point of the application
func main() {

	// Make a GET request to the API endpoint
	resp, err := http.Get(URL)
	if err != nil {
		// Handle any error that occurs while making the request
		fmt.Println("Error:", err)
		return
	}
	// Ensure the response body is closed once the function exits
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Handle any error that occurs while reading the response body
		fmt.Println("Error reading response body:", err)
		return
	}

	// Initialize a variable of type Todo to hold the unmarshaled JSON data
	var todo Todo

	// Unmarshal the JSON response into the Todo struct
	err = json.Unmarshal(body, &todo)
	if err != nil {
		// Handle any error that occurs during the unmarshaling process
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the todo item details in a formatted manner
	fmt.Printf("Todo:\nUserID: %d\nID: %d\nTitle: %s\nCompleted: %v\n", todo.UserID, todo.ID, todo.Title, todo.Completed)
}
