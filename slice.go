package kutils

func SliceHas[T comparable](lst []T, elem T) bool {
	for _, t := range lst {
		if t == elem {
			return true
		}
	}

	return false
}

func SliceMap[T any, R comparable](lst []T, fn func(T) R) []R {
	ret := make([]R, len(lst))
	for i, elem := range lst {
		ret[i] = fn(elem)
	}
	return ret
}

func SliceMapUnique[T any, R comparable](lst []T, fn func(T) R) []R {
	ret := make([]R, 0)
	set := make(map[R]struct{})
	for _, elem := range lst {
		r := fn(elem)
		if _, ok := set[r]; !ok {
			ret = append(ret, r)
			set[r] = struct{}{}
		}
	}
	return ret
}

func SliceMapUniqueWithKeyFn[T, R any, K comparable](lst []T, fn func(T) R, keyFn func(R) K) []R {
	ret := make([]R, 0)
	set := make(map[K]struct{})
	for _, elem := range lst {
		r := fn(elem)
		key := keyFn(r)
		if _, ok := set[key]; !ok {
			ret = append(ret, r)
			set[key] = struct{}{}
		}
	}
	return ret
}

func SliceMapUniques[T any, R comparable](lst []T, fn func(T) []R) []R {
	ret := make([]R, 0)
	set := make(map[R]struct{})
	for _, elem := range lst {
		for _, r := range fn(elem) {
			if _, ok := set[r]; !ok {
				ret = append(ret, r)
				set[r] = struct{}{}
			}
		}
	}
	return ret
}

func SliceMapUniquesWithKeyFn[T, R any, K comparable](lst []T, fn func(T) []R, keyFn func(R) K) []R {
	ret := make([]R, 0)
	set := make(map[K]struct{})
	for _, elem := range lst {
		for _, r := range fn(elem) {
			key := keyFn(r)
			if _, ok := set[key]; !ok {
				ret = append(ret, r)
				set[key] = struct{}{}
			}
		}
	}
	return ret
}

func Chunk[T any](lst []T, chunkSize int) [][]T {
	lstLen := len(lst)
	if lstLen == 0 || chunkSize < 1 {
		return nil
	}

	chunks := make([][]T, 0, (lstLen-1)/chunkSize+1)
	for start, end := 0, 0; start < lstLen; start = end {
		if end = start + chunkSize; end > len(lst) {
			end = len(lst)
		}
		chunks = append(chunks, lst[start:end])
	}

	return chunks
}

// FilterInPlace filters a slice in-place, i.e. it will use the original slice to avoid allocation.
func FilterInPlace[T any](lst []T, filters ...func(elem T) bool) []T {
	if len(lst) == 0 || len(filters) == 0 {
		return lst
	}

	filteredLst := lst[:0]

outer:
	for _, elem := range lst {
		for _, filter := range filters {
			if !filter(elem) {
				continue outer
			}
		}

		filteredLst = append(filteredLst, elem)
	}

	return filteredLst
}
