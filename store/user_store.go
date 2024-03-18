package store

import (
	"bytes"
	"cache-api/models"
	"cache-api/repo"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type UserStoreRepo struct {
	repos *repo.Repos
}

func NewUserStoreRepo(db *repo.Repos) *UserStoreRepo {
	return &UserStoreRepo{
		repos: db,
	}
}

type UserStore interface {
	SeedUserStore(jsonDataPath string) error
	ListUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
}

// Seed database with user data from json file specified in path
// json data should follow shape
//
//	{
//	  "Username": "string"
//	  "Email": "string"
//	  "Age": "number",
//	}
func (s *UserStoreRepo) SeedUserStore(jsonDataPath string) error {
	file, err := os.Open(jsonDataPath)
	if err != nil {
		return fmt.Errorf("error opening JSON file: %w", err)
	}
	defer file.Close()

	// Read the file contents
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, file)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %w", err)
	}

	// Unmarshal JSON data into array of User structs
	var users []models.User
	err = json.Unmarshal(buffer.Bytes(), &users)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON data: %w", err)
	}

	for _, user := range users {
		_, err := s.CreateUser(user)
		if err != nil {
			return fmt.Errorf("error creating user: %w", err)
		}
	}
	log.Print("Seeding Complete !!")
	return nil
}

func (data *UserStoreRepo) ListUsers() ([]models.User, error) {
	var users []models.User
	result := data.repos.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (s *UserStoreRepo) CreateUser(user models.User) (models.User, error) {
	result := s.repos.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, fmt.Errorf("error creating user: %w", result.Error)
	}
	return user, nil
}
