package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

type Store struct {
	db *redis.Client
}

func New() *Store {
	var redisClient *redis.Client

	RedisURL := os.Getenv("REDIS_URL")
	RedisHost := os.Getenv("REDIS_HOST")
	RedisPort := os.Getenv("REDIS_PORT")

	if RedisHost == "" {
		RedisHost = "localhost"
	}
	if RedisPort == "" {
		RedisPort = "6379"
	}

	if RedisURL != "" {
		opt, err := redis.ParseURL(RedisURL)
		if err != nil {
			panic(fmt.Sprintf("Invalid Redis URL: %v", err))
		}
		redisClient = redis.NewClient(opt)
	} else {
		redisClient = redis.NewClient(&redis.Options{
			Addr: RedisHost + ":" + RedisPort,
		})
	}

	// Test Redis connection
	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	fmt.Printf("Connected to Redis: %s\n", pong)

	return &Store{
		db: redisClient,
	}
}

func (s *Store) SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	statusCmd := s.db.WithContext(ctx).Set(key, value, duration)
	if statusCmd.Err() != nil {
		return statusCmd.Err()
	}
	return nil
}

func (s *Store) GetX(ctx context.Context, key string) (interface{}, error) {
	resp := s.db.WithContext(ctx).Get(key)
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	return resp.Val(), nil
}

func (s *Store) DelX(ctx context.Context, key string) error {
	statusCmd := s.db.WithContext(ctx).Del(key)
	if statusCmd.Err() != nil {
		return statusCmd.Err()
	}
	return nil
}
