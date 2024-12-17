package token

type AddressSymbol struct {
	Symbol  string `json:"symbol,omitempty"`
	Address string `json:"address,omitempty"`
}

var MapUSDQuoteTokenByChainId = map[string]*AddressSymbol{
	"42161":  {Symbol: "USDT", Address: "0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9"},
	"43114":  {Symbol: "USDT.e", Address: "0xc7198437980c041c805a1edcba50c1ce5db95118"},
	"8453":   {Symbol: "USDC", Address: "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913"},
	"1":      {Symbol: "USDT", Address: "0xdac17f958d2ee523a2206206994597c13d831ec7"},
	"137":    {Symbol: "USDT", Address: "0xc2132d05d31c914a87c6611c10748aeb04b58e8f"},
	"5000":   {Symbol: "USDT", Address: "0x201eba5cc46d216ce6dc03f6a759e8e766e956ae"},
	"324":    {Symbol: "USDT", Address: "0x493257fd37edb34451f62edf8d2a0c418852ba4c"},
	"81457":  {Symbol: "USDB", Address: "0x4300000000000000000000000000000000000003"},
	"250":    {Symbol: "axlUSDC", Address: "0x1b6382dbdea11d97f24495c9a90b7c88469134a4"},
	"59144":  {Symbol: "USDT", Address: "0xa219439258ca9da29e9cc4ce5596924745e12b93"},
	"10":     {Symbol: "USDT", Address: "0x94b008aa00579c1307b0ef2c499ad98a8ce58e58"},
	"1101":   {Symbol: "USDT", Address: "0x1e4a5963abfd975d8c9021ce480b42188849d41d"},
	"534352": {Symbol: "USDT", Address: "0xf55bec9cafdbe8730f096aa55dad6d22d44099df"},
	"196":    {Symbol: "USDT", Address: "0x1e4a5963abfd975d8c9021ce480b42188849d41d"},
	"56":     {Symbol: "USDT", Address: "0x55d398326f99059ff775485246999027b3197955"},
}

func GetQuoteByChainId(chainId string) *AddressSymbol {
	value, ok := MapUSDQuoteTokenByChainId[chainId]
	if !ok {
		return nil
	}
	return value
}
