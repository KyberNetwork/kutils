package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(redisURL string) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})
	return &RedisCache{client: client}
}

func (r *RedisCache) Set(key string, value interface{}, ttl time.Duration) error {
	// Marshal the value to JSON
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}
	return r.client.Set(context.Background(), key, jsonData, ttl).Err()
}

func (r *RedisCache) Get(key string, result any) error {
	value, err := r.client.Get(context.Background(), key).Result()
	if errors.Is(err, redis.Nil) {
		return fmt.Errorf("key not found")
	}
	err = json.Unmarshal([]byte(value), result)
	return err
}

func (r *RedisCache) Del(key string) error {
	return r.client.Del(context.Background(), key).Err()
}
