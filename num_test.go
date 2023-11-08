package kutils

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

type (
	it int32
	ut uint32
)

func BenchmarkItoa(b *testing.B) {
	i := it(-1234)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Itoa(i)
		}
	})
}

func BenchmarkUtoa(b *testing.B) {
	i := ut(1234)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Utoa(i)
		}
	})
}

func TestItoa(t *testing.T) {
	testItoa(t, "int", -1, "-1")
	testItoa(t, "int8", int8(-0), "0")
	testItoa(t, "int16", int16(-3), "-3")
	testItoa(t, "int32", int32(-2147483648), "-2147483648")
	testItoa(t, "int32 alias", it(2147483647), "2147483647")
	testItoa(t, "int64", int64(-9223372036854775808), "-9223372036854775808")
}

func testItoa[T constraints.Signed](t *testing.T, name string, i T, want string) {
	t.Run(name, func(t *testing.T) {
		if got := Itoa(i); got != want {
			t.Errorf("Itoa() = %v, want %v", got, want)
		}
	})
}

func TestUtoa(t *testing.T) {
	testUtoa(t, "uint", uint(1), "1")
	testUtoa(t, "uint8", uint8(2), "2")
	testUtoa(t, "uint16", uint16(3), "3")
	testUtoa(t, "uint32", uint32(4), "4")
	testUtoa(t, "uint32 alias", ut(4294967295), "4294967295")
	testUtoa(t, "uint64", uint64(18446744073709551615), "18446744073709551615")
}

func testUtoa[T constraints.Unsigned](t *testing.T, name string, i T, want string) {
	t.Run(name, func(t *testing.T) {
		if got := Utoa(i); got != want {
			t.Errorf("Utoa() = %v, want %v", got, want)
		}
	})
}

func TestAtoi(t *testing.T) {
	testAtoi(t, "int", "-1", -1, false)
	testAtoi(t, "int8", "-2", int8(-2), false)
	testAtoi(t, "int8 underflow", "-129", int8(0), true)
	testAtoi(t, "int16", "-3", int16(-3), false)
	testAtoi(t, "int16 overflow", "32768", int16(0), true)
	testAtoi(t, "int32", "2147483647", int32(2147483647), false)
	testAtoi(t, "int32 alias overflow", "2147483648", it(0), true)
	testAtoi(t, "int64", "-9223372036854775808", int64(-9223372036854775808), false)
}

func testAtoi[T constraints.Signed](t *testing.T, name, a string, want T, wantErr bool) {
	t.Run(name, func(t *testing.T) {
		got, err := Atoi[T](a)
		if (err != nil) != wantErr {
			t.Errorf("Atoi() got = %v, error = %v, wantErr %v", got, err, wantErr)
			return
		}
		if got != want {
			t.Errorf("Atoi() got = %v, want %v", got, want)
		}
	})
}

func TestAtou(t *testing.T) {
	testAtou(t, "uint", "0", uint(0), false)
	testAtou(t, "uint8", "255", uint8(255), false)
	testAtou(t, "uint8 underflow", "-1", uint8(0), true)
	testAtou(t, "uint16", "65535", uint16(65535), false)
	testAtou(t, "uint16 overflow", "65536", uint16(0), true)
	testAtou(t, "uint32", "4294967295", uint32(4294967295), false)
	testAtou(t, "uint32 alias overflow", "4294967296", ut(0), true)
	testAtou(t, "uint64", "18446744073709551615", uint64(18446744073709551615), false)
}

func testAtou[T constraints.Unsigned](t *testing.T, name, a string, want T, wantErr bool) {
	t.Run(name, func(t *testing.T) {
		got, err := Atou[T](a)
		if (err != nil) != wantErr {
			t.Errorf("Atou() got = %v, error = %v, wantErr %v", got, err, wantErr)
			return
		}
		if got != want {
			t.Errorf("Atou() got = %v, want %v", got, want)
		}
	})
}

func TestMin(t *testing.T) {
	testMin(t, "int", -1, 2, -1)
	testMin(t, "int8", int8(-0), int8(0), int8(0))
	testMin(t, "int16", int16(16), int16(-3), int16(-3))
	testMin(t, "int32", int32(-2147483648), int32(2147483647), int32(-2147483648))
	testMin(t, "int32 alias", it(2147483647), it(-2147483648), it(-2147483648))
	testMin(t, "int64", int64(-9223372036854775808), int64(9), int64(-9223372036854775808))
	testMin(t, "uint", uint(0), uint(0), uint(0))
	testMin(t, "uint8", uint8(0), uint8(255), uint8(0))
	testMin(t, "uint16", uint16(65535), uint16(0), uint16(0))
	testMin(t, "uint32", uint32(4294967295), uint32(4294967294), uint32(4294967294))
	testMin(t, "uint32 alias", uint32(4294967294), uint32(4294967295), uint32(4294967294))
	testMin(t, "uint64", uint64(18446744073709551615), uint64(1), uint64(1))
	testMin(t, "float32", float32(2.001), float32(5.2), float32(2.001))
	testMin(t, "float64", -.2, -9.e99, -9.e99)
	testMin(t, "string", "10", "2", "10")
}

func testMin[T constraints.Ordered](t *testing.T, name string, a, b, want T) {
	t.Run(name, func(t *testing.T) {
		if got := Min(a, b); !reflect.DeepEqual(got, want) {
			t.Errorf("Min() = %v, want %v", got, want)
		}
	})
}

func TestMax(t *testing.T) {
	testMax(t, "int", -1, 2, 2)
	testMax(t, "int8", int8(-0), int8(0), int8(0))
	testMax(t, "int16", int16(16), int16(-3), int16(16))
	testMax(t, "int32", int32(-2147483648), int32(2147483647), int32(2147483647))
	testMax(t, "int32 alias", it(2147483647), it(-2147483648), it(2147483647))
	testMax(t, "int64", int64(-9223372036854775808), int64(9), int64(9))
	testMax(t, "uint", uint(0), uint(0), uint(0))
	testMax(t, "uint8", uint8(0), uint8(255), uint8(255))
	testMax(t, "uint16", uint16(65535), uint16(0), uint16(65535))
	testMax(t, "uint32", uint32(4294967295), uint32(4294967294), uint32(4294967295))
	testMax(t, "uint32 alias", uint32(4294967294), uint32(4294967295), uint32(4294967295))
	testMax(t, "uint64", uint64(18446744073709551615), uint64(1), uint64(18446744073709551615))
	testMax(t, "float32", float32(2.001), float32(5.2), float32(5.2))
	testMax(t, "float64", -.2, -9.e99, -.2)
	testMax(t, "string", "10", "2", "2")
}

func testMax[T constraints.Ordered](t *testing.T, name string, a, b, want T) {
	t.Run(name, func(t *testing.T) {
		if got := Max(a, b); !reflect.DeepEqual(got, want) {
			t.Errorf("Max() = %v, want %v", got, want)
		}
	})
}

func TestAbs(t *testing.T) {
	testAbs(t, "int", -1, 1, false)
	testAbs(t, "int8", int8(-0), int8(0), false)
	testAbs(t, "int16", int16(32767), int16(32767), false)
	testAbs(t, "int32", int32(2147483647), int32(2147483647), false)
	testAbs(t, "int32 overflow", int32(-2147483648), int32(0), true)
	testAbs(t, "int32 alias", it(-2147483647), it(2147483647), false)
	testAbs(t, "int64", int64(9), int64(9), false)
	testAbs(t, "int64 overflow", int64(-9223372036854775808), int64(0), true)
	testAbs(t, "float32", float32(5.2e32), float32(5.2e32), false)
	testAbs(t, "float64", -9.2e-92, 9.2e-92, false)
}

func testAbs[T constraints.Signed | constraints.Float](t *testing.T, name string, a, want T, wantPanic bool) {
	t.Run(name, func(t *testing.T) {
		var got T
		defer func() {
			if r := recover(); (r != nil) != wantPanic {
				t.Errorf("Abs() got = %v, panic = %v, wantPanic %v", got, r, wantPanic)
			}
		}()
		if got = Abs(a); !reflect.DeepEqual(got, want) {
			t.Errorf("Abs() = %v, want %v", got, want)
		}
	})
}
