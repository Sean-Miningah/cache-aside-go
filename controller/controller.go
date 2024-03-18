package controller

import (
	"cache-api/store"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	store *store.UserStoreRepo
}

func NewUserController(store *store.UserStoreRepo) *UserController {
	return &UserController{
		store: store,
	}
}

func (con *UserController) PingHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Pong!")
}

func (con *UserController) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := con.store.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert users slice to JSON
	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Write JSON response
	w.Write(usersJSON)
}

// func (con *UserController) CacheHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	ctx := context.Background()

// 	err := con.repos.Cache.Set(ctx, "foo", "bar", 0).Err()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Create a map to hold the key-value pair
// 	result := map[string]string{
// 		"key": "foo",
// 		// "value": val,
// 	}

// 	// Convert the map to JSON
// 	jsonResult, err := json.Marshal(result)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the Content-Type header to indicate JSON response
// 	w.Header().Set("Content-Type", "application/json")

// 	// Write the JSON respo
// 	w.Write(jsonResult)
// }
