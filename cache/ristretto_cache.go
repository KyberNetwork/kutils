package cache

import (
	"fmt"
	"reflect"
	"time"

	"github.com/dgraph-io/ristretto"
)

type RistrettoCache struct {
	cache *ristretto.Cache
}

func NewRistrettoCacheDefault() (*RistrettoCache, error) {
	return NewRistrettoCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
}

func NewRistrettoCache(config *ristretto.Config) (*RistrettoCache, error) {
	cache, err := ristretto.NewCache(config)
	if err != nil {
		return nil, err
	}
	return &RistrettoCache{cache: cache}, nil
}

func (r *RistrettoCache) Del(key string) error {
	r.cache.Del(key)
	return nil
}

func (r *RistrettoCache) Set(key string, value interface{}, ttl time.Duration) error {
	// Marshal the value to JSON
	ok := r.cache.SetWithTTL(key, value, 1, ttl)
	if !ok {
		return fmt.Errorf("could not set key: %s", key)
	}
	r.cache.Wait()
	return nil
}

func (r *RistrettoCache) Get(key string, result any) error {
	// Get value from cache
	value, found := r.cache.Get(key)
	if !found {
		return fmt.Errorf("key not found")
	}

	return assignValue(value, result)
}

func assignValue(value interface{}, result any) error {
	resultVal := reflect.ValueOf(result)

	// Validate result parameter
	if err := validateResult(resultVal); err != nil {
		return err
	}

	elem := resultVal.Elem()
	valueVal := reflect.ValueOf(value)

	// Initialize if needed
	if !elem.IsValid() {
		newVal := reflect.New(elem.Type())
		elem.Set(newVal.Elem())
	}

	// Align value and target types
	valueVal = alignTypes(elem, valueVal)

	// Set the value
	if !valueVal.Type().AssignableTo(elem.Type()) {
		return fmt.Errorf("type mismatch: cannot assign %v to %v", valueVal.Type(), elem.Type())
	}

	elem.Set(valueVal)
	return nil
}

func validateResult(resultVal reflect.Value) error {
	if resultVal.Kind() != reflect.Ptr {
		return fmt.Errorf("result must be a pointer")
	}
	if resultVal.IsNil() {
		return fmt.Errorf("result pointer cannot be nil")
	}
	return nil
}
func alignTypes(target, value reflect.Value) reflect.Value {
	// Handle pointer dereferencing
	if target.Kind() == reflect.Ptr {
		if value.Kind() != reflect.Ptr {
			ptr := reflect.New(value.Type())
			ptr.Elem().Set(value)
			return ptr
		}
	} else if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// Handle map type conversion
	if target.Kind() == reflect.Map && value.Kind() == reflect.Map {
		newMap := reflect.MakeMap(target.Type())
		iter := value.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()

			// Handle interface{} values in the map
			if v.Kind() == reflect.Interface {
				v = v.Elem()
			}

			// Handle slices within maps
			if v.Kind() == reflect.Slice {
				sliceType := value.Type() // Create the correct slice type
				newSlice := reflect.MakeSlice(sliceType, v.Len(), v.Cap())

				for i := 0; i < v.Len(); i++ {
					elem := v.Index(i)
					if elem.Kind() == reflect.Interface {
						elem = elem.Elem()
					}
					newSlice.Index(i).Set(elem)
				}
				v = newSlice
			}

			newMap.SetMapIndex(k, v)
		}
		return newMap
	}

	// Handle slice type conversion
	if target.Kind() == reflect.Slice && value.Kind() == reflect.Slice {
		sliceType := value.Type()
		newSlice := reflect.MakeSlice(sliceType, value.Len(), value.Cap())
		for i := 0; i < value.Len(); i++ {
			elem := value.Index(i)
			if elem.Kind() == reflect.Interface {
				elem = elem.Elem()
			}
			newSlice.Index(i).Set(elem)
		}
		return newSlice
	}

	return value
}
