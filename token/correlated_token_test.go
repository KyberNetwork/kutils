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
	ethereum := token.MapCorrelatedTokens["ethereum"]
	require.Len(t, results, len(ethereum))
	require.Greater(t, len(results["usd"]), 8)
	require.Len(t, results["eth"], len(ethereum["eth"]))
	require.Len(t, results["btc"], len(ethereum["btc"]))

	allTokens := token.GetAllTokenByGroup()
	require.Greater(t, len(allTokens), 5)
}
