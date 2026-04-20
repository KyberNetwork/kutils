package token_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/KyberNetwork/kutils/token"
	"github.com/stretchr/testify/require"
)

func TestGetStableTokensByChainID(t *testing.T) {
	results := token.GetStableTokensByChainID(123)
	require.Len(t, results, 0)

	results = token.GetStableTokensByChainID(56)
	require.Len(t, results, 4)
	fmt.Printf("%s", strings.Join(results, "','"))
}

func TestGetDefaultStable(t *testing.T) {
	require.Equal(t, "0xdac17f958d2ee523a2206206994597c13d831ec7", token.GetDefaultStable(1))
}
