package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define the struct to match the API response
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const URL = "https://jsonplaceholder.typicode.com/todos/"

// GetTodo makes a GET request to fetch a Todo item by its ID
func GetTodo(id string) {
	reqURL := URL + id
	// Making a GET request
	resp, err := http.Get(reqURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Reading the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Unmarshal the response into the Todo struct
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the Todo struct in a formatted way
	fmt.Printf("Todo:\nUserID: %d\nID: %d\nTitle: %s\nCompleted: %v\n", todo.UserID, todo.ID, todo.Title, todo.Completed)
}
