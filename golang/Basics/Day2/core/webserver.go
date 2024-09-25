package main

import (
	"io"
	"net/http"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}
func main() {
	http.HandleFunc("/", sayHelloWorld)
	http.ListenAndServe(":8080", nil)
}
