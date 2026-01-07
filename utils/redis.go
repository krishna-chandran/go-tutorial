package utils

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    "log"
    "time"
)

var (
    RedisClient *redis.Client
    ctx         = context.Background()
)

// Initialize Redis connection
func InitRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     "redis:6379", // If using docker network inside compose, service name is "redis:6379"
        Password: "",               // no password
        DB:       0,                // default DB
    })

    pong, err := RedisClient.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }
    fmt.Println("Connected to Redis:", pong)
}

// Set cache with TTL
func SetCache(key string, value string, ttl time.Duration) error {
    return RedisClient.Set(ctx, key, value, ttl).Err()
}

// Get cache
func GetCache(key string) (string, error) {
    val, err := RedisClient.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", fmt.Errorf("key does not exist")
    }
    return val, err
}

// Delete cache
func DeleteCache(key string) error {
    return RedisClient.Del(ctx, key).Err()
}
