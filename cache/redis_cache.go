package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	RedisAddresses string
	Password       string
	MasterName     string
}

type RedisCache struct {
	client redis.UniversalClient
}

func NewRedisCache(config *RedisConfig) *RedisCache {
	addrs := strings.Split(config.RedisAddresses, ",")
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Password:   config.Password,
		Addrs:      addrs,
		MasterName: config.MasterName,
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
