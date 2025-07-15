package models

import "time"

type Prices struct {
	Symbol    string    `json:"symbol,omitempty"`
	Price     float64   `json:"price,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Exchange  string    `json:"exchange,omitempty"`
}

type Sourse struct {
	SourseChan chan Prices
	Addr       string
}
