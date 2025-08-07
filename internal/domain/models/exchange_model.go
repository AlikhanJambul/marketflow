package models

import "time"

type Prices struct {
	Symbol    string    `json:"symbol,omitempty"`
	Price     float64   `json:"price,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Exchange  string    `json:"exchange,omitempty"`
}

type PriceStats struct {
	Exchange  string    `json:"exchange,omitempty"`
	Pair      string    `json:"pair_name,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Average   float64   `json:"average,omitempty"`
	Min       float64   `json:"lowest,omitempty"`
	Max       float64   `json:"highest,omitempty"`
}

type LatestPrice struct {
	Exchange  string    `json:"exchange,omitempty"`
	Pair      string    `json:"pair_name,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Price     float64   `json:"price,omitempty"`
}

//type Sourse struct {
//	Addr       string
//}
