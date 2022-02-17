package util

var supportedCurrencies = []string{
	"USD",
	"GBP",
	"EUR",
	"RMB",
}

func IsSupportedCurrency(currency string) bool {
	for _, sc := range supportedCurrencies {
		if sc == currency {
			return true
		}
	}
	return false
}
