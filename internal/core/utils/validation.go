package utils

type Validation struct {
	Exchanges map[string]struct{}
	Symbols   map[string]struct{}
}

func NewValidation() *Validation {
	return &Validation{
		Exchanges: map[string]struct{}{
			"exchange1": {},
			"exchange2": {},
			"exchange3": {},
			"":          {},
		},
		Symbols: map[string]struct{}{
			"BTCUSDT":  {},
			"ETHUSDT":  {},
			"DOGEUSDT": {},
			"TONUSDT":  {},
			"SOLUSDT":  {},
		},
	}
}

func (v *Validation) CheckSymbol(symbol string) bool {
	if _, ok := v.Symbols[symbol]; ok {
		return true
	}
	return false
}

func (v *Validation) CheckExchange(exchange string) bool {
	if _, ok := v.Exchanges[exchange]; ok {
		return true
	}
	return false
}

func (v *Validation) CheckAll(symbol, exchange string) bool {
	if _, ok := v.Symbols[symbol]; !ok {
		return false
	}

	if _, ok := v.Exchanges[exchange]; !ok {
		return false
	}

	return true
}

func LastPrice(prices []float64) float64 {
	if len(prices) == 0 {
		return 0.0
	}

	return prices[len(prices)-1]
}

func CheckDuration(input string) (string, bool) {
	durations := map[string]string{
		"1s":  "1 second",
		"3s":  "3 seconds",
		"5s":  "5 seconds",
		"10s": "10 seconds",
		"30s": "30 seconds",
		"1m":  "1 minute",
		"3m":  "3 minutes",
		"5m":  "5 minutes",
	}

	d, ok := durations[input]
	return d, ok
}
