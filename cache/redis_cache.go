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

// RedisConfig contains all configuration of redis
type RedisConfig struct {
	Addresses        string
	MasterName       string
	DBNumber         int
	Username         string
	Password         string
	SentinelUsername string
	SentinelPassword string
	Prefix           string
	Separator        string
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	RouteRandomly    bool
	ReplicaOnly      bool
}

type RedisCache struct {
	client redis.UniversalClient
}

func NewRedisCache(cfg *RedisConfig) *RedisCache {
	addrs := strings.Split(cfg.Addresses, ",")
	if cfg.MasterName != "" {
		client := redis.NewFailoverClusterClient(&redis.FailoverOptions{
			MasterName:       cfg.MasterName,
			SentinelAddrs:    addrs,
			DB:               cfg.DBNumber,
			Username:         cfg.Username,
			Password:         cfg.Password,
			SentinelUsername: cfg.SentinelUsername,
			SentinelPassword: cfg.SentinelPassword,
			ReadTimeout:      cfg.ReadTimeout,
			WriteTimeout:     cfg.WriteTimeout,
			RouteRandomly:    cfg.RouteRandomly,
			ReplicaOnly:      cfg.ReplicaOnly,
		})
		return &RedisCache{client: client}
	}
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            addrs,
		MasterName:       cfg.MasterName,
		DB:               cfg.DBNumber,
		Username:         cfg.Username,
		Password:         cfg.Password,
		SentinelUsername: cfg.SentinelUsername,
		SentinelPassword: cfg.SentinelPassword,
		ReadTimeout:      cfg.ReadTimeout,
		WriteTimeout:     cfg.WriteTimeout,
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
