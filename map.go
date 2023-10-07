package kutils

func Map[T any, K comparable, V any](lst []T, keyFn func(T) K, valFn func(T) V) map[K]V {
	res := make(map[K]V, len(lst))
	for _, elem := range lst {
		res[keyFn(elem)] = valFn(elem)
	}
	return res
}

func MapMulti[T any, K comparable, V any](lst []T, keyFn func(T) K, valFn func(T) V) map[K][]V {
	res := make(map[K][]V)
	for _, elem := range lst {
		res[keyFn(elem)] = append(res[keyFn(elem)], valFn(elem))
	}
	return res
}

func MapKey[T any, K comparable](lst []T, keyFn func(T) K) map[K]T {
	res := make(map[K]T, len(lst))
	for _, elem := range lst {
		res[keyFn(elem)] = elem
	}
	return res
}

func MapKeyMulti[T any, K comparable](lst []T, keyFn func(T) K) map[K][]T {
	res := make(map[K][]T)
	for _, elem := range lst {
		res[keyFn(elem)] = append(res[keyFn(elem)], elem)
	}
	return res
}
