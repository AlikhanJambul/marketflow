package models

import "time"

type Exchange struct {
	Symbol    string
	Value     float64
	Timestamp time.Time
	Source    string
}
