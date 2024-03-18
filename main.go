package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"cache-api/repo"
	"cache-api/store"
)

const Port = 8080

func main() {
	var (
		migrate bool
		seed    bool
	)

	flag.BoolVar(&migrate, "migrate", false, "run database migration")
	flag.BoolVar(&seed, "seed", false, "seed database with appropriate data")
	flag.Parse()

	config := repo.Config{
		DBHost:        "localhost",
		DBPort:        "5432",
		DBUser:        "postgres",
		DBName:        "redis-cache-project",
		DBPassword:    "postgres",
		CacheHost:     "localhost",
		CacheUser:     "default",
		CachePassword: "my-password",
	}

	repos, err := repo.NewRepos(config)
	if err != nil {
		log.Fatal(err)
	}

	if migrate {
		repos.MakeMigrations()
	}

	userStore := store.NewUserStoreRepo(repos)

	if seed {
		userStore.SeedUserStore("./test-data/users.json")
	}

	// Simple Handler function for '/ping' endpoint
	pingHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!")
	}

	cacheHandler := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		err := repos.Cache.Set(ctx, "foo", "bar", 0).Err()
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
