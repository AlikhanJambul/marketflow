package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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
	re := regexp.MustCompile(`(\d+)([hms])`)
	matches := re.FindAllStringSubmatch(input, -1)

	if len(matches) == 0 {
		return "", false
	}

	var parts []string
	for _, match := range matches {
		valueStr, unit := match[1], match[2]
		value, _ := strconv.Atoi(valueStr)

		switch unit {
		case "h":
			if value == 1 {
				parts = append(parts, fmt.Sprintf("%d hour", value))
			} else {
				parts = append(parts, fmt.Sprintf("%d hours", value))
			}
		case "m":
			if value == 1 {
				parts = append(parts, fmt.Sprintf("%d minute", value))
			} else {
				parts = append(parts, fmt.Sprintf("%d minutes", value))
			}
		case "s":
			if value == 1 {
				parts = append(parts, fmt.Sprintf("%d second", value))
			} else {
				parts = append(parts, fmt.Sprintf("%d seconds", value))
			}
		}
	}

	return strings.Join(parts, " "), true
}
