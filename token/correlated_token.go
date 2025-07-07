package token

import (
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
	"strings"
)

var MapCorrelatedTokens = map[string]map[string]map[string]string{
	"ethereum": {
		"usd": {
			"USDC":  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			"USDT":  "0xdac17f958d2ee523a2206206994597c13d831ec7",
			"DAI":   "0x6b175474e89094c44da98b954eedeac495271d0f",
			"MAI":   "0x8D6CeBD76f18E1558D4DB88138e2DeFB3909fAD6",
			"BOB":   "0xB0B195aEFA3650A6908f15CdaC7D92F8a5791B0B",
			"MIM":   "0x99D8a9C45b2ecA8864373A26D1459e3Dff1e17F3",
			"USDe":  "0x4c9edd5852cd905f086c759e8383e09bff1e68b3",
			"sUSDe": "0x9D39A5DE30e57443BfF2A8307A4256c8797A3497",
			"USD1":  "0x8d0D000Ee44948FC98c9B98A4FA4921476f08B0d",
		},
		"eth": {
			"WETH":          "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			"ezETH":         "0xbf5495Efe5DB9ce00f80364C8B423567e58d2110",
			"weETH":         "0xcd5fe23c85820f7b72d0926fc9b05b43e359b7ee",
			"rsETH":         "0xa1290d69c65a6fe4df752f95823fae25cb99e5a7",
			"stETH":         "0xae7ab96520de3a18e5e111b5eaab095312d7fe84",
			"wstETH":        "0x7f39c581f595b53c5cb19bd0b3f8da6c935e2ca0",
			"pufETH":        "0xd9a442856c234a39a81a089c06451ebaa4306a72",
			"rswETH":        "0xFAe103DC9cf190eD75350761e95403b7b8aFa6c0",
			"mETH":          "0xd5F7838F5C461fefF7FE49ea5ebaF7728bB0ADfa",
			"frxETH":        "0x5E8422345238F34275888049021821E8E08CAa1f",
			"sfrxETH":       "0xac3E018457B222d93114458476f3E3416Abbe38F",
			"swETH":         "0xf951E335afb289353dc249e82926178EaC7DEd78",
			"ethx":          "0xA35b1B31Ce002FBF2058D22F30f95D405200A15b",
			"oETH":          "0x856c4Efb76C1D1AE02e20CEB03A2A6a08b0b8dC3",
			"primeETH":      "0x6ef3D766Dfe02Dc4bF04aAe9122EB9A0Ded25615",
			"bedrockUniETH": "0xF1376bceF0f78459C0Ed0ba5ddce976F1ddF51F4",
			"wbETH":         "0xa2E3356610840701BDf5611a53974510Ae27E2e1",
			"rETH":          "0xae78736Cd615f374D3085123A210448E74Fc6393",
			"eETH":          "0x35fA164735182de50811E8e2E824cFb9B6118ac2",
			"ETH":           "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		},
		"btc": {
			"WBTC": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
			"sBTC": "0xfE18be6b3Bd88A2D2A7f928d00292E7a9963CfC6",
		},
	},
	"polygon": {
		"usd": {
			"USDT":   "0xc2132d05d31c914a87c6611c10748aeb04b58e8f",
			"USDC.e": "0x2791bca1f2de4661ed88a30c99a7a9449aa84174",
			"USDC":   "0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359",
			"DAI":    "0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063",
			"MAI":    "0xa3Fa99A148fA48D14Ed51d610c367C61876997F1",
			"BOB":    "0xB0B195aEFA3650A6908f15CdaC7D92F8a5791B0B",
			"MIM":    "0x49a0400587A7F65072c87c4910449fDcC5c47242",
		},
		"matic": {
			"WMATIC":  "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270",
			"MATICX":  "0xfa68FB4628DFF1028CFEc22b4162FCcd0d45efb6",
			"stMATIC": "0x3a58a54c066fdc0f2d55fc9c89f0415c92ebf3c4",
			"rMatic":  "0x9f28e2455f9FFcFac9EBD6084853417362bc5dBb",
			"POL":     "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		},
		"eth": {
			"WETH":   "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619",
			"wstETH": "0x03b54a6e9a984069379fae1a4fc4dbae93b3bccd",
		},
	},
	"bsc": {
		"usd": {
			"USDT":    "0x55d398326f99059ff775485246999027b3197955",
			"USDC":    "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d",
			"DAI":     "0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3",
			"BUSD":    "0xe9e7cea3dedca5984780bafc599bd69add087d56",
			"MAI":     "0x3F56e0c36d275367b8C502090EDF38289b3dEa0d",
			"BOB":     "0xB0B195aEFA3650A6908f15CdaC7D92F8a5791B0B",
			"MIM":     "0xfE19F0B51438fd612f6FD59C1dbB3eA319f433Ba",
			"axlUSDC": "0x4268b8f0b87b6eae5d897996e6b845ddbd99adf3",
		},
		"eth": {
			"wbETH": "0xa2E3356610840701BDf5611a53974510Ae27E2e1",
		},
	},
	"avalanche": {
		"usd": {
			"USDT":   "0x9702230A8Ea53601f5cD2dc00fDBc13d4dF4A8c7",
			"USDC":   "0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E",
			"DAI.e":  "0xd586E7F844cEa2F87f50152665BCbc2C279D8d70",
			"USDT.e": "0xc7198437980c041c805A1EDcbA50c1Ce5db95118",
			"USDC.e": "0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664",
			"MAI":    "0x5c49b268c9841AFF1Cc3B0a418ff5c3442eE3F3b",
			"YUSD":   "0x111111111111ed1D73f860F57b2798b683f2d325",
			"MIM":    "0x130966628846BFd36ff31a822705796e8cb8C18D",
		},
		"avax": {
			"WAVAX": "0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7",
			"sAVAX": "0x2b2C81e08f1Af8835a78Bb2A90AE924ACE0eA4bE",
			"AVAX":  "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		},
	},
	"arbitrum": {
		"usd": {
			"USDT":    "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9",
			"USDC":    "0xaf88d065e77c8cC2239327C5EDb3A432268e5831",
			"DAI":     "0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1",
			"USDC.e":  "0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8",
			"MAI":     "0x3F56e0c36d275367b8C502090EDF38289b3dEa0d",
			"MIM":     "0xFEa7a6a0B346362BF88A9e4A88416B77a57D6c2A",
			"fUSDC":   "0x4cfa50b7ce747e2d61724fcac57f24b748ff2b2a",
			"axlUSDC": "0xeb466342c4d449bc9f53a865d5cb90586f405215",
		},
		"eth": {
			"WETH":    "0x82af49447d8a07e3bd95bd0d56f35241523fbab1",
			"ezETH":   "0x2416092f143378750bb29b79ed961ab195cceea5",
			"weETH":   "0x35751007a407ca6feffe80b3cb397736d2cf4dbe",
			"wstETH":  "0x5979D7b546E38E414F7E9822514be443A4800529",
			"rETH":    "0xec70dcb4a1efa46b8f2d97c310c9c4790ba5ffa8",
			"sfrxETH": "0x95ab45875cffdba1e5f451b950bc2e42c0053f39",
			"frxETH":  "0x178412e79c25968a32e89b11f63b33f733770c2a",
		},
		"btc": {
			"tBTC":  "0x6c84a8f1c29108F47a79964b5Fe888D4f4D0dE40",
			"BTC.b": "0x2297aebd383787a160dd0d9f71508148769342e3",
			"WBTC":  "0x2f2a2543b76a4166549f7aab2e75bef0aefc5b0f",
		},
	},
	"optimism": {
		"usd": {
			"USDT":    "0x94b008aa00579c1307b0ef2c499ad98a8ce58e58",
			"USDC.e":  "0x7f5c764cbc14f9669b88837ca1490cca17c31607",
			"USDC":    "0x0b2c639c533813f4aa9d7837caf62653d097ff85",
			"DAI":     "0xda10009cbd5d07dd0cecc66161fc93d7c9000da1",
			"MAI":     "0xdFA46478F9e5EA86d57387849598dbFB2e964b02",
			"BOB":     "0xB0B195aEFA3650A6908f15CdaC7D92F8a5791B0B",
			"sUSD":    "0x8c6f28f2f1a3c87f0f938b96d27520d9751ec8d9",
			"crvUSD":  "0xc52d7f23a2e460248db6ee192cb23dd12bddcbf6",
			"LUSD":    "0xc40f949f8a4e094d1b49a23ea9241d289b7b2819",
			"USD+":    "0x73cb180bf0521828d8849bc8cf2b920918e23032",
			"axlUSDC": "0xeb466342c4d449bc9f53a865d5cb90586f405215",
		},
		"eth": {
			"WETH":   "0x4200000000000000000000000000000000000006",
			"wstETH": "0x1F32b1c2345538c0c6f582fCB022739c4A194Ebb",
			"rETH":   "0x9bcef72be871e61ed4fbbc7630889bee758eb81d",
			"ezETH":  "0x2416092f143378750bb29b79ed961ab195cceea5",
			"sETH":   "0xe405de8f52ba7559f9df3c368500b6e6ae6cee49",
			"cbETH":  "0xaddb6a0412de1ba0f936dcaeb8aaa24578dcf3b2",
			"frxETH": "0x6806411765af15bddd26f8f544a34cc40cb9838b",
			"ETH":    "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		},
		"btc": {
			"sBTC": "0x298b9b95708152ff6968aafd889c6586e9169f1d",
			"WBTC": "0x68f180fcce6836688e9084f035309e29bf0a2095",
			"tBTC": "0x6c84a8f1c29108f47a79964b5fe888d4f4d0de40",
		},
	},
	"base": {
		"usd": {
			"USDC":    "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",
			"USDbC":   "0xd9aAEc86B65D86f6A7B5B1b0c42FFA531710b6CA",
			"DAI":     "0x50c5725949A6F0c72E6C4a641F24049A917DB0Cb",
			"MAI":     "0xbf1aeA8670D2528E08334083616dD9C5F3B087aE",
			"USD+":    "0xb79dd08ea68a908a97220c76d19a6aa9cbde4376",
			"axlUSDC": "0xeb466342c4d449bc9f53a865d5cb90586f405215",
			"EURC":    "0x60a3e35cc302bfa44cb288bc5a4f316fdb1adb42",
		},
		"eth": {
			"WETH":   "0x4200000000000000000000000000000000000006",
			"wstETH": "0xc1cba3fcea344f92d9239c08c0568f6f2f0ee452",
			"cbETH":  "0x2ae3f1ec7f1f5012cfeab0185bfc7aa3cf0dec22",
			"ETH":    "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		},
	},
	"linea": {
		"usd": {
			"USDT": "0xA219439258ca9da29E9Cc4cE5596924745e12B93",
			"USDC": "0x176211869cA2b568f2A7D4EE941E073a821EE1ff",
			"DAI":  "0x4AF15ec2A0BD43Db75dd04E62FAA3B8EF36b00d5",
			"MAI":  "0xf3B001D64C656e30a62fbaacA003B1336b4ce12A",
		},
		"eth": {
			"WETH":   "0xe5d7c2a44ffddf6b295a15c148167daaaf5cf34f",
			"wstETH": "0xb5bedd42000b71fdde22d3ee8a79bd49a568fc8f",
			"ezETH":  "0x2416092f143378750bb29b79ed961ab195cceea5",
			"inETH":  "0x5a7a183b6b44dc4ec2e3d2ef43f98c5152b1d76d",
			"uniETH": "0x15eefe5b297136b8712291b632404b66a8ef4d25",
			"ETH":    "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		},
	},
	"scroll": {
		"usd": {
			"USDT": "0xf55BEC9cafDbE8730f096Aa55dad6D22d44099Df",
			"USDC": "0x06eFdBFf2a14a7c8E15944D1F4A48F9F95F663A4",
			"LUSD": "0xeDEAbc3A1e7D21fE835FFA6f83a710c70BB1a051",
			"DAI":  "0xcA77eB3fEFe3725Dc33bccB54eDEFc3D9f764f97",
			"USDe": "0x5d3a1ff2b6bab83b63cd9ad0787074081a52ef34",
		},
		"eth": {
			"WETH":   "0x5300000000000000000000000000000000000004",
			"rETH":   "0x53878b874283351d26d206fa512aece1bef6c0dd",
			"pufETH": "0xc4d46E8402F476F269c379677C99F18E22Ea030e",
			"weETH":  "0x01f0a31698c4d065659b9bdc21b3610292a1c506",
			"wstETH": "0xf610a9dfb7c89644979b4a0f27063e9e7d7cda32",
			"ETH":    "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		},
	},
	"blast": {
		"usd": {
			"USDB": "0x4300000000000000000000000000000000000003",
			"MIM":  "0x76da31d7c9cbeae102aff34d3398bc450c8374c1",
		},
		"eth": {
			"WETH":  "0x4300000000000000000000000000000000000004",
			"ezETH": "0x2416092f143378750bb29b79ed961ab195cceea5",
			"nETH":  "0xce971282fAAc9faBcF121944956da7142cccC855",
			"weETH": "0x04c0599ae5a44757c0af6f9ec3b93da8976c150a",
			"ETH":   "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		},
	},
	"mantle": {},
	"zksync": {
		"usd": {
			"USDT":   "0x493257fD37EDB34451f62EDf8D2a0C418852bA4C",
			"USDC":   "0x1d17CBcF0D6D143135aE902365D2E5e2A16538D4",
			"USDC.e": "0x3355df6D4c9C3035724Fd0e3914dE96A5a83aaf4",
			"DAI":    "0x4B9eb6c0b6ea15176BBF62841C6B2A8a398cb656",
			"crvUSD": "0x43cD37CC4B9EC54833c8aC362Dd55E58bFd62b86",
			"USD+":   "0x8E86e46278518EFc1C5CEd245cBA2C7e3ef11557",
			"LUSD":   "0x503234F203fC7Eb888EEC8513210612a43Cf6115",
			"hUSDC":  "0xaf08a9d918f16332F22cf8Dc9ABE9D9E14DdcbC2",
			"cBUSD":  "0x2039bb4116B4EFc145Ec4f0e2eA75012D6C0f181",
		},
		"eth": {
			"WETH":  "0x5aea5775959fbc2557cc8789bc1bf90a239d9a91",
			"rETH":  "0x32Fd44bB869620C0EF993754c8a00Be67C464806",
			"cbETH": "0x75Af292c1c9a37b3EA2E6041168B4E48875b9ED5",
			"ETH":   "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		},
	},
	"fantom": {},
	"polygon-zkevm": {
		"eth": {
			"WETH":   "0x4F9A0e7FD2Bf6067db6994CF12E4495Df938E6e9",
			"rETH":   "0xb23c20efce6e24acca0cef9b7b7aa196b84ec942",
			"frxETH": "0xCf7eceE185f19e2E970a301eE37F93536ed66179",
			"ETH":    "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		},
	},
	"sonic": {
		"usd": {
			"USDC": "0x29219dd400f2Bf60E5a23d13Be72B486D4038894",
			"USDT": "0x6047828dc181963ba44974801ff68e538da5eaf9",
			"USD+": "0x53e24706D6642CA495498557415b1af7A025D8Da",
		},
	},
	"berachain": {
		"usd": {
			"USDC":  "0x549943e04f40284185054145c6E4e9568C1D3241",
			"byUSD": "0x688e72142674041f8f6Af4c808a4045cA1D6aC82",
			"USDT":  "0x779Ded0c9e1022225f8E0630b35a9b54bE713736",
		},
		"eth": {
			"WETH":   "0x4F9A0e7FD2Bf6067db6994CF12E4495Df938E6e9",
			"rETH":   "0xb23c20efce6e24acca0cef9b7b7aa196b84ec942",
			"frxETH": "0xCf7eceE185f19e2E970a301eE37F93536ed66179",
		},
	},
	"ronin": {},
	"unichain": {
		"usd": {
			"USDC":  "0x078D782b760474a361dDA0AF3839290b0EF57AD6",
			"USDT0": "0x9151434b16b9763660705744891fA906F660EcC5",
			"USDT":  "0x588CE4F028D8e7B53B687865d6A67b3A54C75518",
			"DAI":   "0x20CAb320A855b39F724131C69424240519573f81",
		},
	},
}

func GetTokensByGroup(chainId uint) map[string][]string {
	results := make(map[string][]string)
	chainName, err := valueobject.ToString(valueobject.ChainID(chainId))
	if err != nil {
		return results
	}

	mapGroupTokens, ok := MapCorrelatedTokens[chainName]
	if !ok {
		return results
	}
	for group := range mapGroupTokens {
		tokenAddress := make([]string, 0, len(mapGroupTokens[group]))
		for token := range mapGroupTokens[group] {
			tokenAddress = append(tokenAddress, strings.ToLower(mapGroupTokens[group][token]))
		}
		results[group] = tokenAddress
	}
	return results
}

func GetAllTokenByGroup() map[string][]string {
	results := make(map[string][]string)
	for chain := range MapCorrelatedTokens {
		for group := range MapCorrelatedTokens[chain] {
			if _, ok := results[group]; !ok {
				results[group] = make([]string, 0)
			}
			tokenAddress := results[group]
			for token := range MapCorrelatedTokens[chain][group] {
				tokenAddress = append(tokenAddress, strings.ToLower(MapCorrelatedTokens[chain][group][token]))
			}
			results[group] = tokenAddress
		}
	}
	return results
}
