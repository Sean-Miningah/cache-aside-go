package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"cache-api/controller"
	"cache-api/repo"
	"cache-api/store"

	"github.com/julienschmidt/httprouter"
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

	userservice := controller.NewUserController(userStore)

	router := httprouter.New()
	router.GET("/ping", userservice.PingHandler)
	router.GET("/users", userservice.GetUsers)

	// Start the HTTP server on the specified port
	fmt.Printf("Server listening on :%d \n", Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", Port), router)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
