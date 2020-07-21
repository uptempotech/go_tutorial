package main

import (
	"fmt"
	"net/http"
)

// Create an HTTP server

// http.ResponseWriter assembles the servers response and writes to the client.
// http.Request is the client's request
func handler(w http.ResponseWriter, r *http.Request) {
	// Writes to the client
	fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth\n")
}

func main() {

	// Calls for function handler output to match the directory /
	http.HandleFunc("/", handler)

	// Calls for function handler2 output to match directory /earth
	http.HandleFunc("/earth", handler2)

	fmt.Println("Starting HTTP server on port 8080")
	// Listen to port 8080 and handle requests
	http.ListenAndServe(":8080", nil)
}
