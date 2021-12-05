package currencies

type Currency struct {
	Name                  string
	Type                  CurrencyType
	UrlPrefix             string
	Decimals              uint
	PendingTimeoutSeconds int // Mempool timeout in seconds
	VerifiedTimeoutBlocks int // Confirmation timeout in # of blocks
	TokenAddress          string
}

type CurrencyType int

const (
	Coin = iota
	EthToken
)

var pTimeout = 600
var ethereumVTimeout = 30 // in blocks

var Currencies = []string{}
var Chains = []string{}
var EthTokens = []string{}
var CurrencyData = map[string]Currency{
	"btc": {
		Name:                  "Bitcoin",
		Type:                  Coin,
		UrlPrefix:             "bitcoin",
		Decimals:              8,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: 30,
	},
	"eth": {
		Name:                  "Ethereum",
		Type:                  Coin,
		UrlPrefix:             "ethereum",
		Decimals:              18,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
	},
	"bch": {
		Name:                  "Bitcoin Cash",
		Type:                  Coin,
		UrlPrefix:             "bitcoincash",
		Decimals:              8,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: 25,
	},
	"ltc": {
		Name:                  "Litecoin",
		Type:                  Coin,
		UrlPrefix:             "litecoin",
		Decimals:              8,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: 25,
	},
	"doge": {
		Name:                  "Dogecoin",
		Type:                  Coin,
		UrlPrefix:             "dogecoin",
		Decimals:              8,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: 25,
	},
	"dai": {
		Name:                  "Dai",
		Type:                  EthToken,
		UrlPrefix:             "dai",
		Decimals:              18,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
		TokenAddress:          "0x6b175474e89094c44da98b954eedeac495271d0f",
	},
	"usdt": {
		Name:                  "Tether",
		Type:                  EthToken,
		UrlPrefix:             "tether",
		Decimals:              6,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
		TokenAddress:          "0xdac17f958d2ee523a2206206994597c13d831ec7",
	},
	"usdc": {
		Name:                  "USDC",
		Type:                  EthToken,
		UrlPrefix:             "usdc",
		Decimals:              18,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
		TokenAddress:          "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	},
	"ampl": {
		Name:                  "Ampleforth",
		Type:                  EthToken,
		UrlPrefix:             "ampl",
		Decimals:              9,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
		TokenAddress:          "0xd46ba6d942050d489dbd938a2c909a5d5039a161",
	},
}

func SupportAllEthTokens() {
	for k, v := range CurrencyData {
		if v.Type == EthToken {
			EthTokens = append(EthTokens, k)
		}
	}
}

var EthTokensOld = []string{
	"dai",
	"usdt",
	"usdc",
	"ust",
	"ampl",
}
