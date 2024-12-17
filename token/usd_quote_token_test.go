package token

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetStableTokensByChainID(t *testing.T) {
	require.Equal(t, "0xdac17f958d2ee523a2206206994597c13d831ec7", GetQuoteByChainId("1").Address)
	require.Nil(t, GetQuoteByChainId("11"))
}
