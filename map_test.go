package kutils

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	testMap[int, int, string](t, "int to map[abs]string",
		[]int{-1, 2, -3, 2}, Abs[int], Itoa[int],
		map[int]string{1: "-1", 2: "2", 3: "-3"})
	testMap[uint, string, uint](t, "uint to map[string]cappedIncr",
		[]uint{1, 2, 9, 3}, Utoa[uint], func(u uint) uint { return Min(9, u+1) },
		map[string]uint{"1": 2, "2": 3, "9": 9, "3": 4})
}

func testMap[T any, K comparable, V any](t *testing.T, name string, lst []T, keyFn func(T) K, valFn func(T) V,
	want map[K]V) {
	t.Run(name, func(t *testing.T) {
		if got := Map(lst, keyFn, valFn); !reflect.DeepEqual(got, want) {
			t.Errorf("Map() = %v, want %v", got, want)
		}
	})
}

func TestMapMulti(t *testing.T) {
	testMapMulti[int, int, string](t, "int to map[abs][]string",
		[]int{-1, 2, -3, -2}, Abs[int], Itoa[int],
		map[int][]string{1: {"-1"}, 2: {"2", "-2"}, 3: {"-3"}})
	testMapMulti[uint, string, uint](t, "uint to map[string][]cappedIncr",
		[]uint{1, 2, 9, 3}, Utoa[uint], func(u uint) uint { return Min(9, u+1) },
		map[string][]uint{"1": {2}, "2": {3}, "9": {9}, "3": {4}})
}

func testMapMulti[T any, K comparable, V any](t *testing.T, name string, lst []T, keyFn func(T) K, valFn func(T) V,
	want map[K][]V) {
	t.Run(name, func(t *testing.T) {
		if got := MapMulti(lst, keyFn, valFn); !reflect.DeepEqual(got, want) {
			t.Errorf("MapMulti() = %v, want %v", got, want)
		}
	})
}

func TestMapKey(t *testing.T) {
	testMapKey[int, int](t, "int to map[abs]",
		[]int{-1, 2, -3, 2}, Abs[int],
		map[int]int{1: -1, 2: 2, 3: -3})
	testMapKey[uint, string](t, "uint to map[string]",
		[]uint{1, 2, 9, 3}, Utoa[uint],
		map[string]uint{"1": 1, "2": 2, "9": 9, "3": 3})
}

func testMapKey[T any, K comparable](t *testing.T, name string, lst []T, keyFn func(T) K, want map[K]T) {
	t.Run(name, func(t *testing.T) {
		if got := MapKey(lst, keyFn); !reflect.DeepEqual(got, want) {
			t.Errorf("MapKey() = %v, want %v", got, want)
		}
	})
}

func TestMapKeyMulti(t *testing.T) {
	testMapKeyMulti[int, int](t, "int to map[abs]",
		[]int{-1, 2, -3, -2}, Abs[int],
		map[int][]int{1: {-1}, 2: {2, -2}, 3: {-3}})
	testMapKeyMulti[uint, string](t, "uint to map[string]",
		[]uint{1, 2, 9, 3}, Utoa[uint],
		map[string][]uint{"1": {1}, "2": {2}, "9": {9}, "3": {3}})
}

func testMapKeyMulti[T any, K comparable](t *testing.T, name string, lst []T, keyFn func(T) K, want map[K][]T) {
	t.Run(name, func(t *testing.T) {
		if got := MapKeyMulti(lst, keyFn); !reflect.DeepEqual(got, want) {
			t.Errorf("MapKeyMulti() = %v, want %v", got, want)
		}
	})
}
