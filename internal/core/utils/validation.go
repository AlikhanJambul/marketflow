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
