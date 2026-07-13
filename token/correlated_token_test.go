package token_test

import (
	"testing"

	"github.com/KyberNetwork/kutils/token"
	"github.com/stretchr/testify/require"
)

func TestGetTokensByGroup(t *testing.T) {
	results := token.GetTokensByGroup(123)
	require.Len(t, results, 0)

	results = token.GetTokensByGroup(1)
	require.Len(t, results, 4)
	require.Greater(t, len(results["usd"]), 8)
	require.Len(t, results["eth"], 18)
	require.Len(t, results["btc"], 4)

	allTokens := token.GetAllTokenByGroup()
	require.Greater(t, len(allTokens), 5)
}
