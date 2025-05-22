package token

import (
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
	"strings"
)

var MapStableTokens = map[string]map[string]string{
	"ethereum": {
		"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48": "USDC",
		"0xdac17f958d2ee523a2206206994597c13d831ec7": "USDT",
		"0x6b175474e89094c44da98b954eedeac495271d0f": "DAI",
	},
	"polygon": {
		"0xc2132d05d31c914a87c6611c10748aeb04b58e8f": "USDT",
		"0x2791bca1f2de4661ed88a30c99a7a9449aa84174": "USDC.e",
		"0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359": "USDC",
		"0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063": "DAI",
	},
	"bsc": {
		"0x55d398326f99059ff775485246999027b3197955": "USDT",
		"0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d": "USDC",
		"0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3": "DAI",
		"0xe9e7cea3dedca5984780bafc599bd69add087d56": "BUSD",
	},
	"avalanche": {
		"0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7": "USDT",
		"0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e": "USDC",
		"0xd586e7f844cea2f87f50152665bcbc2c279d8d70": "DAI.e",
		"0xc7198437980c041c805a1edcba50c1ce5db95118": "USDT.e",
		"0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664": "USDC.e",
	},
	"arbitrum": {
		"0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9": "USDT",
		"0xaf88d065e77c8cc2239327c5edb3a432268e5831": "USDC",
		"0xda10009cbd5d07dd0cecc66161fc93d7c9000da1": "DAI",
		"0xff970a61a04b1ca14834a43f5de4533ebddb5cc8": "USDC.e",
	},
	"optimism": {
		"0x94b008aa00579c1307b0ef2c499ad98a8ce58e58": "USDT",
		"0x7f5c764cbc14f9669b88837ca1490cca17c31607": "USDC",
		"0xda10009cbd5d07dd0cecc66161fc93d7c9000da1": "DAI",
		"0x7F5c764cBc14f9669B88837ca1490cCa17c31607": "USDT.e",
	},
	"base": {
		"0x833589fcd6edb6e08f4c7c32d4f71b54bda02913": "USDC",
		"0xd9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca": "USDbC",
		"0x50c5725949a6f0c72e6c4a641f24049a917db0cb": "DAI",
	},
	"linea": {
		"0xa219439258ca9da29e9cc4ce5596924745e12b93": "USDT",
		"0x176211869ca2b568f2a7d4ee941e073a821ee1ff": "USDC",
		"0x4af15ec2a0bd43db75dd04e62faa3b8ef36b00d5": "DAI",
	},
	"scroll": {
		"0xf55bec9cafdbe8730f096aa55dad6d22d44099df": "USDT",
		"0x06efdbff2a14a7c8e15944d1f4a48f9f95f663a4": "USDC",
	},
	"blast": {
		"0x4300000000000000000000000000000000000003": "USDB",
	},
	"mantle": {
		"0x201eba5cc46d216ce6dc03f6a759e8e766e956ae": "USDT",
		"0x09bc4e0d864854c6afb6eb9a9cdf58ac190d0df9": "USDC",
	},
	"zksync": {
		"0x493257fd37edb34451f62edf8d2a0c418852ba4c": "USDT",
		"0x3355df6d4c9c3035724fd0e3914de96a5a83aaf4": "USDC.e",
		"0x1d17cbcf0d6d143135ae902365d2e5e2a16538d4": "USDC",
		"0x4B9eb6c0b6ea15176BBF62841C6B2A8a398cb656": "DAI",
	},
	"fantom": {
		"0x049d68029688eabf473097a2fc38ef61633a3c7a": "fUSDT",
		"0x04068da6c83afcfa0e13ba15a6696662335d5b75": "USDC",
		"0x8d11ec38a3eb5e956b052f67da8bdc9bef8abf3e": "DAI",
		"0xcc1b99dDAc1a33c201a742A1851662E87BC7f22C": "USDT",
	},
	"polygon-zkevm": {
		"0xa8ce8aee21bc2a48a5ef670afcc9274c7bbbc035": "USDC",
		"0x1e4a5963abfd975d8c9021ce480b42188849d41d": "USDT",
		"0xc5015b9d9161dca7e18e32f6f25c4ad850731fd4": "DAI",
		"0x37eaa0ef3549a5bb7d431be78a3d99bd360d19e5": "USDC.e",
	},
	"sonic": {
		"0x29219dd400f2Bf60E5a23d13Be72B486D4038894": "USDC",
		"0x6047828dc181963ba44974801ff68e538da5eaf9": "USDT",
	},
	"berachain": {
		"0x549943e04f40284185054145c6E4e9568C1D3241": "USDC",
		"0x688e72142674041f8f6Af4c808a4045cA1D6aC82": "byUSD",
		"0x779Ded0c9e1022225f8E0630b35a9b54bE713736": "USDT",
	},
	"ronin": {
		"0x0b7007c13325c48911f73a2dad5fa5dcbf808adc": "USDC",
	},
	"unichain": {
		"0x078D782b760474a361dDA0AF3839290b0EF57AD6": "USDC",
		"0x9151434b16b9763660705744891fA906F660EcC5": "USDT0",
		"0x588CE4F028D8e7B53B687865d6A67b3A54C75518": "USDT",
		"0x20CAb320A855b39F724131C69424240519573f81": "DAI",
	},
}

func GetStableTokensByChainID(chainId uint) []string {
	chainName, err := valueobject.ToString(valueobject.ChainID(chainId))
	if err != nil {
		return []string{}
	}

	mapTokens, ok := MapStableTokens[chainName]
	if !ok {
		return []string{}
	}
	listToken := make([]string, 0, len(mapTokens))
	for key := range mapTokens {
		listToken = append(listToken, strings.ToLower(key))
	}
	return listToken
}
