package currencies

type Currency struct {
	Name                  string
	Type                  CurrencyType
	UrlPrefix             string
	Decimals              uint
	PendingTimeoutSeconds int // Mempool timeout in seconds
	VerifiedTimeoutBlocks int // Confirmation timeout in # of blocks
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
	},
	"usdt": {
		Name:                  "Tether",
		Type:                  EthToken,
		UrlPrefix:             "tether",
		Decimals:              6,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
	},
	"usdc": {
		Name:                  "USDC",
		Type:                  EthToken,
		UrlPrefix:             "usdc",
		Decimals:              18,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
	},
	"ampl": {
		Name:                  "Ampleforth",
		Type:                  EthToken,
		UrlPrefix:             "ampl",
		Decimals:              9,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
	},
}

var EthTokensOld = []string{
	"dai",
	"usdt",
	"usdc",
	"ust",
	"ampl",
}
