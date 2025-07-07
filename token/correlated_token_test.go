package token_test

import (
	"github.com/KyberNetwork/kutils/token"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetTokensByGroup(t *testing.T) {
	results := token.GetTokensByGroup(123)
	require.Len(t, results, 0)

	results = token.GetTokensByGroup(1)
	require.Len(t, results, 3)
	require.Greater(t, len(results["usd"]), 8)
	require.Len(t, results["eth"], 20)
	require.Len(t, results["btc"], 2)

	allTokens := token.GetAllTokenByGroup()
	require.Len(t, allTokens, 5)
}
