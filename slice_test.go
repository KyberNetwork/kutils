package kutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceHas(t *testing.T) {
	testSliceHas(t, "int", []int{1, 2, 3}, 2, true)
	testSliceHas(t, "uint", []uint{1, 2, 3}, 4, false)
	testSliceHas(t, "string", []string{"a", "b"}, "C", false)
	type st string
	testSliceHas(t, "string alias", []st{"a", "b"}, "a", true)
	ptr := &struct{ I int }{1}
	testSliceHas(t, "ptr", []*struct{ I int }{{}, ptr}, &struct{ I int }{}, false)
	testSliceHas(t, "ptr same", []*struct{ I int }{ptr}, ptr, true)
	ch := make(chan int)
	testSliceHas(t, "chan", []chan int{ch}, (chan int)(nil), false)
	testSliceHas(t, "chan same", []chan int{ch}, ch, true)
	testSliceHas(t, "array", [][2]int{{1, 2}}, [2]int{1, 2}, true)
	testSliceHas(t, "struct", []struct{ I int }{{1}}, *ptr, true)
	testSliceHas(t, "interface", []any{1}, any(1), true)
}

func testSliceHas[T comparable](t *testing.T, name string, lst []T, elem T, want bool) {
	t.Run(name, func(t *testing.T) {
		assert.Equalf(t, want, SliceHas(lst, elem), "SliceHas(%v, %v)", lst, elem)
	})
}

func TestSliceMap(t *testing.T) {
	testSliceMap(t, "Itoa", []int{1, 2, 3}, Itoa[int], []string{"1", "2", "3"})
	testSliceMap(t, "div 2", []int{1, 2, 3}, func(i int) int { return i / 2 }, []int{0, 1, 1})
}

func testSliceMap[T any, R comparable](t *testing.T, name string, lst []T, fn func(T) R, want []R) {
	t.Run(name, func(t *testing.T) {
		assert.Equalf(t, want, SliceMap(lst, fn), "SliceMap(%v, fn)", lst)
	})
}

func TestSliceMapUnique(t *testing.T) {
	testSliceMapUnique(t, "Itoa", []int{1, 2, 3}, Itoa[int], []string{"1", "2", "3"})
	testSliceMapUnique(t, "div 2", []int{1, 2, 3}, func(i int) int { return i / 2 }, []int{0, 1})
}

func testSliceMapUnique[T any, R comparable](t *testing.T, name string, lst []T, fn func(T) R, want []R) {
	t.Run(name, func(t *testing.T) {
		assert.Equalf(t, want, SliceMapUnique(lst, fn), "SliceMapUnique(%v, fn)", lst)
	})
}

func TestSliceMapUniqueWithKeyFn(t *testing.T) {
	testSliceMapUniqueWithKeyFn(t, "div 2", []int{1, 2, 3, 2, 1}, func(i int) int { return i + 1 },
		func(i int) int { return i / 2 }, []int{2, 4})
}

func testSliceMapUniqueWithKeyFn[T any, R any, K comparable](t *testing.T, name string, lst []T, fn func(T) R,
	keyFn func(R) K, want []R) {
	t.Run(name, func(t *testing.T) {
		assert.Equalf(t, want, SliceMapUniqueWithKeyFn(lst, fn, keyFn), "SliceMapUniqueWithKeyFn(%v, fn, keyFn)", lst)
	})
}

func TestSliceMapUniques(t *testing.T) {
	testSliceMapUniques(t, "div 2 and div 3", []int{1, 2, 3, 4, 9}, func(i int) []int { return []int{i / 2, i / 3} },
		[]int{0, 1, 2, 4, 3})
}

func testSliceMapUniques[T any, R comparable](t *testing.T, name string, lst []T, fn func(T) []R, want []R) {
	t.Run(name, func(t *testing.T) {
		assert.Equalf(t, want, SliceMapUniques(lst, fn), "SliceMapUniques(%v, fn)", lst)
	})
}

func TestSliceMapUniquesWithKeyFn(t *testing.T) {
	testSliceMapUniquesWithKeyFn(t, "add 1 and 2, unique by div 3", []int{1, 2, 4, 10},
		func(i int) []int { return []int{i + 1, i + 2} },
		func(i int) int { return i / 3 }, []int{2, 3, 6, 11, 12})
}

func testSliceMapUniquesWithKeyFn[T any, R any, K comparable](t *testing.T, name string, lst []T, fn func(T) []R,
	keyFn func(R) K, want []R) {
	t.Run(name, func(t *testing.T) {
		assert.Equalf(t, want, SliceMapUniquesWithKeyFn(lst, fn, keyFn), "SliceMapUniquesWithKeyFn(%v, fn, keyFn)", lst)
	})
}

func TestChunk(t *testing.T) {
	testChunk(t, "happy", []int{1, 2, 3, 5, 8}, 3, [][]int{{1, 2, 3}, {5, 8}})
	testChunk(t, "chunk size divides", []int{1, 2, 3, 5}, 2, [][]int{{1, 2}, {3, 5}})
	testChunk(t, "chunk size big", []int{1, 2, 3}, 4, [][]int{{1, 2, 3}})
	testChunk(t, "empty lst", ([]int)(nil), 3, nil)
	testChunk(t, "invalid chunk size", []int{1, 2, 3, 5, 8}, 0, nil)
}

func testChunk[T any](t *testing.T, name string, lst []T, chunkSize int, want [][]T) {
	t.Run(name, func(t *testing.T) {
		assert.Equalf(t, want, Chunk(lst, chunkSize), "Chunk(%v, %v)", lst, chunkSize)
	})
}

func TestFilterInPlace(t *testing.T) {
	testFilterInPlace(t, "happy", []int{1, 2, 3, 5, 8}, []func(int) bool{func(i int) bool { return i%2 < 1 }},
		[]int{2, 8})
	testFilterInPlace(t, "empty lst", ([]int)(nil), []func(int) bool{func(i int) bool { return i%2 < 1 }}, nil)
	testFilterInPlace(t, "no filters", []int{1, 2, 3, 5}, nil, []int{1, 2, 3, 5})
	testFilterInPlace(t, "multiple filters", []int{1, 2, 3, 5, 8},
		[]func(int) bool{func(i int) bool { return i > 1 }, func(i int) bool { return i%2 > 0 }}, []int{3, 5})
}

func testFilterInPlace[T any](t *testing.T, name string, lst []T, filters []func(T) bool, want []T) {
	t.Run(name, func(t *testing.T) {
		assert.Equalf(t, want, FilterInPlace(lst, filters...), "FilterInPlace(%v, filters...)", lst)
	})
}
