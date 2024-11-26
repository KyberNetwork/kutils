package cache

import (
	"github.com/dgraph-io/ristretto"
	"time"
)

type Cache interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string, result any) error
	Del(key string) error
}

type CfgCache struct {
	*ristretto.Config
	Type  string
	Redis *RedisConfig
}

func NewCache(cfg *CfgCache) Cache {
	var cache Cache
	if cfg.Type == "redis" && cfg.Redis != nil {
		cache = NewRedisCache(cfg.Redis)
	} else {
		if cfg.Config == nil {
			cache, _ = NewRistrettoCacheDefault()
		} else {
			cache, _ = NewRistrettoCache(&ristretto.Config{
				NumCounters: cfg.NumCounters,
				MaxCost:     cfg.MaxCost,
				BufferItems: cfg.BufferItems,
			})
		}
	}

	return cache
}
