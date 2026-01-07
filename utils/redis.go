package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"time"
)

var (
	RedisClient *redis.Client
	ctx         = context.Background()
)

// InitRedis initializes Redis client
func InitRedis() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "redis:6379" // default for Docker Compose
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	pong, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis:", pong)
}

// SetCache sets a value in Redis with TTL
func SetCache(key string, value string, ttl time.Duration) error {
	return RedisClient.Set(ctx, key, value, ttl).Err()
}

// GetCache fetches a value from Redis
func GetCache(key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key does not exist")
	}
	return val, err
}

// DeleteCache removes a key from Redis
func DeleteCache(key string) error {
	return RedisClient.Del(ctx, key).Err()
}

// CacheFetch generic helper for fetching from cache or DB
func CacheFetch[T any](key string, ttl time.Duration, fetchFunc func() (T, error)) (T, error) {
	var result T

	cached, err := GetCache(key)
	if err == nil {
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result, nil
		}
	}

	// Cache miss
	result, err = fetchFunc()
	if err != nil {
		return result, err
	}

	data, _ := json.Marshal(result)
	SetCache(key, string(data), ttl)

	return result, nil
}
