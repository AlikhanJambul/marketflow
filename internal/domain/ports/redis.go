package ports

import (
	"context"
	"marketflow/internal/domain/models"
)

type Cache interface {
	Get(string) (models.PriceStats, error)
	Set(string, string, models.PriceStats) error
	SetLatest(prices models.Prices) error
	GetLatest(string) (models.LatestPrice, error)
	Check(context.Context) error
}
