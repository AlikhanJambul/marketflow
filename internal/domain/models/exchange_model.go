package models

import "time"

type Prices struct {
	Symbol    string    `json:"symbol,omitempty"`
	Exchange  string    `json:"exchange,omitempty"`
	Price     float64   `json:"price,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

type PriceStats struct {
	Pair      string    `json:"pair_name,omitempty"`
	Exchange  string    `json:"exchange,omitempty"`
	Average   float64   `json:"average,omitempty"`
	Min       float64   `json:"lowest,omitempty"`
	Max       float64   `json:"highest,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type LatestPrice struct {
	Pair      string    `json:"pair_name,omitempty"`
	Exchange  string    `json:"exchange,omitempty"`
	Price     float64   `json:"latest_price,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
