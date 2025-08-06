package ports

import (
	"context"
	"marketflow/internal/domain/models"
)

type Cache interface {
	Get(string) (models.PriceStats, error)
	Set(string, models.PriceStats) error
	SetLatest(string, string, models.LatestPrice) error
	GetLatest(string) (models.LatestPrice, error)
	Check(context.Context) error
}
