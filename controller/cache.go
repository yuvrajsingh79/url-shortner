package controller

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

// Cache defines an interface for caching URL mappings.
type Cache interface {
	Set(key, value string) error
	Get(key string) (string, bool)
}

// RedisCache is a Redis-based implementation of the Cache interface.
type RedisCache struct {
	client *redis.Client
}

// / NewRedisCache creates a new RedisCache instance.
func NewRedisCache(redisAddr string) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // No password
		DB:       0,  // Default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisCache{client: client}, nil
}

// Set stores a key-value pair in the Redis cache with an expiration time.
func (c *RedisCache) Set(key, value string) error {
	// ctx := context.Background()
	expiration := 24 * time.Hour // Adjust the expiration time as needed.

	err := c.client.Set(key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves a value from the Redis cache by key.
func (c *RedisCache) Get(key string) (string, bool) {
	val, err := c.client.Get(key).Result()
	if err != nil && err == redis.Nil {
		return "", false
	} else if err != nil {
		log.Printf("Redis cache error: %v", err)
		return "", false
	}

	return val, true
}
