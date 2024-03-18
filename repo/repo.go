package repo

import (
	"fmt"
	"log"

	"cache-api/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBName        string
	DBPassword    string
	CacheHost     string
	CacheUser     string
	CachePassword string
}

type Repos struct {
	Cache *redis.Client
	DB    *gorm.DB
}

func NewRepos(cfg Config) (*Repos, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Africa/Nairobi",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBName,
		cfg.DBPassword,
	)
	// sqlDB, err := sql.Open("postgres", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	opt, err := redis.ParseURL(fmt.Sprintf(
		"redis://%s:%s@%s:6379/0",
		cfg.CacheUser, cfg.CachePassword, cfg.CacheHost))
	if err != nil {
		return nil, err
	}

	cache := redis.NewClient(opt)

	return &Repos{
		Cache: cache,
		DB:    db,
	}, nil
}

func (repos Repos) MakeMigrations() error {
	repos.DB.AutoMigrate(
		&models.User{},
	)
	log.Printf("Migration Complete !!")
	return nil
}
