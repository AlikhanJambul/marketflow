package utils

import "time"

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

func CheckDuration(input string) (time.Duration, bool) {
	durations := map[string]time.Duration{
		"1s":  time.Second,
		"3s":  3 * time.Second,
		"5s":  5 * time.Second,
		"10s": 10 * time.Second,
		"30s": 30 * time.Second,
		"1m":  time.Minute,
		"3m":  3 * time.Minute,
		"5m":  5 * time.Minute,
	}

	d, ok := durations[input]
	return d, ok
}
