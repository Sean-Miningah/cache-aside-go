package store

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
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

type Store struct {
	Cache *redis.Client
	DB    *sql.DB
}

func NewStore(cfg Config) (*Store, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBName,
		cfg.DBPassword,
	)
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	// test db connection
	err = sqlDB.Ping()
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

	return &Store{
		Cache: cache,
		DB:    sqlDB,
	}, nil
}
