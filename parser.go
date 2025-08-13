package kutils

import (
	"math/big"
)

const LimitLength = 200

func SafeParseBigInt(s string, base int) (*big.Int, bool) {
	if len(s) > LimitLength {
		return nil, false
	}
	return new(big.Int).SetString(s, base)
}
