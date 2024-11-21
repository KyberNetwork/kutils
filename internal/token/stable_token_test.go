package token_test

import (
	"fmt"
	"github.com/KyberNetwork/kutils/internal/token"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestGetStableTokensByChainID(t *testing.T) {
	results := token.GetStableTokensByChainID(123)
	require.Len(t, results, 0)

	results = token.GetStableTokensByChainID(56)
	require.Len(t, results, 4)
	fmt.Printf("%s", strings.Join(results, "','"))
}
