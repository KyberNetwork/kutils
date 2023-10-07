package kutils

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
)

func Itoa[T constraints.Signed](i T) string {
	return strconv.FormatInt(int64(i), 10)
}

func Utoa[T constraints.Unsigned](u T) string {
	return strconv.FormatUint(uint64(u), 10)
}

func Atoi[T constraints.Signed](a string) (i T, err error) {
	bitSize := 0
	switch any(i).(type) {
	case int:
		bitSize = strconv.IntSize
	case int8:
		bitSize = 8
	case int16:
		bitSize = 16
	case int32:
		bitSize = 32
	case int64:
		bitSize = 64
	}
	if i, err := strconv.ParseInt(a, 10, bitSize); err != nil {
		return 0, err
	} else if bitSize == 0 && int64(T(i)) != i {
		return 0, &strconv.NumError{Func: "ParseInt", Num: a, Err: strconv.ErrRange}
	} else {
		return T(i), nil
	}
}

func Atou[T constraints.Unsigned](a string) (u T, err error) {
	bitSize := 0
	switch any(u).(type) {
	case uint, uintptr:
		bitSize = strconv.IntSize
	case uint8:
		bitSize = 8
	case uint16:
		bitSize = 16
	case uint32:
		bitSize = 32
	case uint64:
		bitSize = 64
	}
	if u, err := strconv.ParseUint(a, 10, bitSize); err != nil {
		return 0, err
	} else if bitSize == 0 && uint64(T(u)) != u {
		return 0, &strconv.NumError{Func: "ParseUint", Num: a, Err: strconv.ErrRange}
	} else {
		return T(u), nil
	}
}

func Min[T constraints.Ordered](a, b T) T {
	if b < a {
		return b
	}
	return a
}

func Max[T constraints.Ordered](a, b T) T {
	if b > a {
		return b
	}
	return a
}

func Abs[T constraints.Signed | constraints.Float](a T) T {
	if a < 0 {
		if a-1 > 0 {
			panic(fmt.Sprintf("Abs result of %v will overflow", a))
		}
		return -a
	}
	return a
}
