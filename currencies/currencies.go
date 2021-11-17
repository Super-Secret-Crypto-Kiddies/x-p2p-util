package currencies

type Currency struct {
	Name                  string
	UrlPrefix             string
	Decimals              uint
	PendingTimeoutSeconds int // Mempool timeout in seconds
	VerifiedTimeoutBlocks int // Confirmation timeout in # of blocks
}

var pTimeout = 600
var ethereumVTimeout = 30 // in blocks

var Currencies = map[string]Currency{
	"btc": {
		Name:                  "Bitcoin",
		UrlPrefix:             "bitcoin",
		Decimals:              8,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: 30,
	},
	"eth": {
		Name:                  "Ethereum",
		UrlPrefix:             "ethereum",
		Decimals:              18,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
	},
	"bch": {
		Name:                  "Bitcoin Cash",
		UrlPrefix:             "bitcoincash",
		Decimals:              8,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: 25,
	},
	"ltc": {
		Name:                  "Litecoin",
		UrlPrefix:             "litecoin",
		Decimals:              8,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: 25,
	},
	"doge": {
		Name:                  "Dogecoin",
		UrlPrefix:             "dogecoin",
		Decimals:              8,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: 25,
	},
	"dai": {
		Name:                  "Dai",
		UrlPrefix:             "dai",
		Decimals:              18,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
	},
	"usdt": {
		Name:                  "Tether",
		UrlPrefix:             "tether",
		Decimals:              6,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
	},
	"usdc": {
		Name:                  "USDC",
		UrlPrefix:             "usdc",
		Decimals:              18,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
	},
	"ampl": {
		Name:                  "Ampleforth",
		UrlPrefix:             "ampl",
		Decimals:              9,
		PendingTimeoutSeconds: pTimeout,
		VerifiedTimeoutBlocks: ethereumVTimeout,
	},
}

var EthTokens = []string{
	"dai",
	"usdt",
	"usdc",
	"ust",
	"ampl",
}
