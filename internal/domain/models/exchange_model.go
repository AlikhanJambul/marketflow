package models

import "time"

type Prices struct {
	PairName  string    `json:"pair_name,omitempty"`
	Value     float64   `json:"value,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Exchange  string    `json:"exchange,omitempty"`
}

type Sourse struct {
	SourseChan chan Prices
	Addr       string
}
