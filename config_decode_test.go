package kutils

import (
	"math"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/KyberNetwork/kutils/internal/json"
)

func Test_DecodeConfig(t *testing.T) {
	type args struct {
		data json.RawMessage
		dest any
	}
	ptr := func(x any) *any { return &x }
	bigInt := func(s string) *big.Int { x, _ := new(big.Int).SetString(s, 10); return x }
	bigFloat := func(s string) *big.Float { x, _ := new(big.Float).SetString(s); return x }
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
		wantVal any
	}{
		{
			"object",
			args{
				data: json.RawMessage(`{"foo": "qhx", "bar": 1}`),
				dest: &struct {
					Foo string
					Bar *int
				}{Foo: "bar", Bar: nil},
			},
			assert.NoError,
			&struct {
				Foo string
				Bar *int
			}{Foo: "qhx", Bar: &[]int{1}[0]},
		},
		{
			"map",
			args{
				data: json.RawMessage(`{"foo": "qhz", "1": "3"}`),
				dest: &map[string]string{"qux": "bar", "1": "2"},
			},
			assert.NoError,
			&map[string]string{"qux": "bar", "foo": "qhz", "1": "3"},
		},
		{
			"array",
			args{
				data: json.RawMessage(`["foo", "bar"]`),
				dest: &[]string{"foo"},
			},
			assert.NoError,
			&[]string{"foo", "bar"},
		},
		{
			"string",
			args{
				data: json.RawMessage(`"bar"`),
				dest: ptr("foo"),
			},
			assert.NoError,
			ptr("bar"),
		},
		{
			"number",
			args{
				data: json.RawMessage("456"),
				dest: ptr(123),
			},
			assert.NoError,
			ptr(456),
		},
		{
			"bool",
			args{
				data: json.RawMessage("true"),
				dest: ptr(false),
			},
			assert.NoError,
			ptr(true),
		},
		{
			"string to duration",
			args{
				data: json.RawMessage(`"1.ms"`),
				dest: ptr(time.Second),
			},
			assert.NoError,
			ptr(time.Millisecond),
		},
		{
			"empty string to duration",
			args{
				data: json.RawMessage(`""`),
				dest: ptr(time.Second),
			},
			assert.NoError,
			ptr(time.Duration(0)),
		},
		{
			"decimal string to duration",
			args{
				data: json.RawMessage(`"-.5h"`),
				dest: ptr(time.Second),
			},
			assert.NoError,
			ptr(-time.Hour / 2),
		},
		{
			"number to duration",
			args{
				data: json.RawMessage("1000000"),
				dest: ptr(time.Second),
			},
			assert.NoError,
			ptr(time.Millisecond),
		},
		{
			"big number to uint64",
			args{
				data: json.RawMessage("18446744073709551615"),
				dest: ptr(uint64(0)),
			},
			assert.NoError,
			ptr(uint64(math.MaxUint64)),
		},
		{
			"big number to float64",
			args{
				data: json.RawMessage("1.7976931348623157E308"),
				dest: ptr(float64(0)),
			},
			assert.NoError,
			ptr(math.MaxFloat64),
		},
		{
			"big number overflowing int64",
			args{
				data: json.RawMessage("18446744073709551617"),
				dest: ptr(int64(0)),
			},
			assert.Error,
			nil,
		},
		{
			"big number to big.Int",
			args{
				data: json.RawMessage("18446744073709551617"),
				dest: ptr(big.Int{}),
			},
			assert.NoError,
			ptr(*bigInt("18446744073709551617")),
		},
		{
			"big number to *big.Int",
			args{
				data: json.RawMessage("18446744073709551617"),
				dest: ptr(new(big.Int)),
			},
			assert.NoError,
			ptr(bigInt("18446744073709551617")),
		},
		{
			"big number string to *big.Int",
			args{
				data: json.RawMessage(`"18446744073709551617123000"`),
				dest: ptr(new(big.Int)),
			},
			assert.NoError,
			ptr(bigInt("18446744073709551617123000")),
		},
		{
			"invalid *big.Int",
			args{
				data: json.RawMessage("1.8076931348623157E-309"),
				dest: ptr(new(big.Int)),
			},
			assert.Error,
			nil,
		},
		{
			"big number overflowing float64",
			args{
				data: json.RawMessage("1.7976931348623157E309"),
				dest: ptr(float64(0)),
			},
			assert.Error,
			nil,
		},
		{
			"big number to big.Float",
			args{
				data: json.RawMessage("1.807693134862315789001E420"),
				dest: ptr(big.Float{}),
			},
			assert.NoError,
			ptr(*bigFloat("1.807693134862315789001E420")),
		},
		{
			"big number to *big.Float",
			args{
				data: json.RawMessage("1.8076931348623157E309"),
				dest: ptr(new(big.Float)),
			},
			assert.NoError,
			ptr(bigFloat("1.8076931348623157E309")),
		},
		{
			"big number string to *big.Float",
			args{
				data: json.RawMessage(`"1.8076931348623157E9876"`),
				dest: ptr(new(big.Float)),
			},
			assert.NoError,
			ptr(bigFloat("1.8076931348623157E9876")),
		},
		{
			"invalid big.Float",
			args{
				data: json.RawMessage(`"E30"`),
				dest: ptr(new(big.Float)),
			},
			assert.Error,
			nil,
		},
		{
			"big number to decimal.Decimal",
			args{
				data: json.RawMessage("18446744073709551617.2"),
				dest: ptr(decimal.Decimal{}),
			},
			assert.NoError,
			ptr(decimal.RequireFromString("18446744073709551617.2")),
		},
		{
			"big number to *decimal.Decimal",
			args{
				data: json.RawMessage("184467440737095516173E1"),
				dest: ptr(ptr(decimal.Decimal{})),
			},
			assert.NoError,
			ptr(ptr(decimal.RequireFromString("184467440737095516173E1"))),
		},
		{
			"invalid decimal.Decimal",
			args{
				data: json.RawMessage("1E"),
				dest: ptr(new(decimal.Decimal)),
			},
			assert.Error,
			nil,
		},
		{
			"string to address",
			args{
				data: json.RawMessage(`"0x0123456789abcdef0123456789abcdef01234567"`),
				dest: ptr(common.Address{}),
			},
			assert.NoError,
			ptr(common.HexToAddress("0x0123456789abcDEF0123456789abCDef01234567")),
		},
		{
			"string to hash",
			args{
				data: json.RawMessage(`"0x0123456789abcDEF0123456789abCDef0123456789abcDEF0123456789abCDef"`),
				dest: ptr(common.Hash{}),
			},
			assert.NoError,
			ptr(common.HexToHash("0x0123456789abcDEF0123456789abCDef0123456789abcDEF0123456789abCDef")),
		},
		{
			"string to json.Unmarshaler",
			args{
				data: json.RawMessage(`"foo"`),
				dest: ptr(json.RawMessage{}),
			},
			assert.NoError,
			ptr(json.RawMessage(`"foo"`)),
		},
		{
			"string to slice",
			args{
				data: json.RawMessage(`"a,b"`),
				dest: &[]string{},
			},
			assert.NoError,
			&[]string{"a", "b"},
		},
		{
			"malformed json",
			args{
				data: json.RawMessage("malformed"),
				dest: ptr("foo"),
			},
			assert.Error,
			nil,
		},
		{
			"invalid dest",
			args{
				data: json.RawMessage("456"),
				dest: 123,
			},
			assert.Error,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				err := DecodeConfig(tt.args.data, tt.args.dest)
				tt.wantErr(t, err, "DecodeConfig(%v, %v)", tt.args.data, tt.args.dest)
				if err == nil {
					assert.Equal(t, tt.wantVal, tt.args.dest)
				}
			})
		})
	}
}

func TestStringUnmarshalHookFunc(t *testing.T) {
	hook := StringUnmarshalHookFunc()
	data, err := hook(reflect.ValueOf("hi"), reflect.ValueOf(""))
	assert.NoError(t, err)
	assert.Equal(t, "hi", data)
}
