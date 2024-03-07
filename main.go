package main

import (
	"fmt"
	"net/http"
)

const Port = 8080

func main() {

	// Simple Handler function for the ping endpoint
	pingHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!")
	}

	// Register the handler for the "/ping" endpoint
	http.HandleFunc("/ping", pingHandler)

	// Start the HTTP server on the specified port
	fmt.Printf("Server listening on :%d \n", Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
