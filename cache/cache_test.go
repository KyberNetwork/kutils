package cache_test

import (
	"testing"
	"time"

	"github.com/KyberNetwork/kutils/cache"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	cacheTypes := []struct {
		name   string
		config *cache.CfgCache
	}{
		{"Ristretto", &cache.CfgCache{
			Type: "ristretto",
		}},
		{"Redis", &cache.CfgCache{
			Type: "ristretto",
			Redis: &cache.RedisConfig{
				Addresses: "localhost:6379",
			},
		}},
	}

	type example struct {
		Name  string
		Age   int
		Value uint64
	}

	for _, ct := range cacheTypes {
		t.Run(ct.name, func(t *testing.T) {
			var err error

			sCache := cache.NewCache(ct.config)
			// Example usage
			key1 := "exampleKey"
			key2 := "exampleKey1"
			input1 := example{
				Name:  t.Name(),
				Value: 1,
			}
			input2 := "demo 123"
			ttl := 5 * time.Minute

			err = sCache.Set(key1, input1, ttl)
			require.NoError(t, err, "Error setting cache")
			err = sCache.Set(key2, input2, ttl)
			require.NoError(t, err, "Error setting cache")
			var res1 *example
			err = sCache.Get(key1, &res1)
			require.NoError(t, err, "Error getting cache")
			require.Equal(t, input1, *res1, "Error expected cache value")

			var res2 string
			err = sCache.Get(key2, &res2)
			require.NoError(t, err, "Error getting cache")
			require.Equal(t, input2, res2, "Error expected cache value")
		})
		t.Run("Pointer Types", func(t *testing.T) {
			sCache := cache.NewCache(ct.config)
			key := "pointerTest"
			input := &example{Name: "pointer", Age: 25, Value: 100}
			err := sCache.Set(key, input, time.Minute)
			require.NoError(t, err)

			var result *example
			err = sCache.Get(key, &result)
			require.NoError(t, err)
			require.Equal(t, input, result)
		})
		t.Run("Slice Types", func(t *testing.T) {
			sCache := cache.NewCache(ct.config)
			key := "sliceTest"
			input := []*example{{Name: "pointer", Age: 25, Value: 100}}
			err := sCache.Set(key, input, time.Minute)
			require.NoError(t, err)

			var result []*example
			err = sCache.Get(key, &result)
			require.NoError(t, err)
			require.Equal(t, input, result)
		})

		t.Run("Non-existent Key", func(t *testing.T) {
			sCache := cache.NewCache(ct.config)
			var result string
			err := sCache.Get("nonexistentKey", &result)
			require.Error(t, err)
			require.Contains(t, err.Error(), "key not found")
		})

		t.Run("Type Mismatch", func(t *testing.T) {
			sCache := cache.NewCache(ct.config)
			key := "typeMismatch"
			err := sCache.Set(key, 42, time.Minute)
			require.NoError(t, err)

			var result string
			err = sCache.Get(key, &result)
			require.Error(t, err)
		})

		t.Run("Nil Pointer", func(t *testing.T) {
			sCache := cache.NewCache(ct.config)
			key := "nilPointer"
			err := sCache.Set(key, "test", time.Minute)
			require.NoError(t, err)

			var result *string
			err = sCache.Get(key, result) // passing nil pointer
			require.Error(t, err)
			require.Contains(t, err.Error(), "nil")
		})

		t.Run("Non-existent Key", func(t *testing.T) {
			sCache := cache.NewCache(ct.config)
			var result string
			err := sCache.Get("nonexistentKey", &result)
			require.Error(t, err)
			require.Contains(t, err.Error(), "key not found")
		})

		t.Run("Type Mismatch", func(t *testing.T) {
			sCache := cache.NewCache(ct.config)
			key := "typeMismatch"
			err := sCache.Set(key, 42, time.Minute)
			require.NoError(t, err)

			var result string
			err = sCache.Get(key, &result)
			require.Error(t, err)
		})

		t.Run("Nil Pointer", func(t *testing.T) {
			sCache := cache.NewCache(ct.config)
			key := "nilPointer"
			err := sCache.Set(key, "test", time.Minute)
			require.NoError(t, err)

			var result *string
			err = sCache.Get(key, result) // passing nil pointer
			require.Error(t, err)
			require.Contains(t, err.Error(), "nil")
		})

		//t.Run("Complex Types", func(t *testing.T) {
		//	sCache := cache.NewCache(ct.config)
		//	key := "complexType"
		//	input := map[string]interface{}{
		//		"name": "test",
		//		"data": []int{1, 2, 3},
		//	}
		//	err := sCache.Set(key, input, time.Minute)
		//	require.NoError(t, err)
		//
		//	var result *map[string]interface{}
		//	err = sCache.Get(key, &result)
		//	require.NoError(t, err)
		//	require.Equal(t, input, *result)
		//})
	}
}

func TestDelKey(t *testing.T) {
	cacheTypes := []struct {
		name   string
		config *cache.CfgCache
	}{
		{"Ristretto", &cache.CfgCache{
			Type: "ristretto",
		}},
		{"Redis", &cache.CfgCache{
			Type: "ristretto",
			Redis: &cache.RedisConfig{
				Addresses: "localhost:6379",
			},
		}},
	}
	for _, ct := range cacheTypes {
		t.Run(ct.name, func(t *testing.T) {
			sCache := cache.NewCache(ct.config)
			key := "test_key"
			value := "123test"
			err := sCache.Set(key, value, time.Minute)
			require.NoError(t, err)
			err = sCache.Del(key)
			require.NoError(t, err)

			var result string
			err = sCache.Get(key, &result)
			require.Error(t, err)
		})
	}
}
