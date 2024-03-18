package main

import (
	store "cache-api/data"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const Port = 8080

func main() {
	config := store.Config{
		DBHost:        "main-db",
		DBPort:        "5432",
		DBUser:        "postgres",
		DBName:        "postgres",
		DBPassword:    "postgres",
		CacheHost:     "web-server-cache",
		CacheUser:     "default",
		CachePassword: "my-password",
	}

	store, err := store.NewStore(config)
	if err != nil {
		log.Fatal(err)
	}

	// Simple Handler function for '/ping' endpoint
	pingHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!")
	}

	cacheHandler := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		err := store.Cache.Set(ctx, "foo", "bar", 0).Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a map to hold the key-value pair
		result := map[string]string{
			"key": "foo",
			// "value": val,
		}

		// Convert the map to JSON
		jsonResult, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to indicate JSON response
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON respo
		w.Write(jsonResult)
	}

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/cache", cacheHandler)

	// Start the HTTP server on the specified port
	fmt.Printf("Server listening on :%d \n", Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
