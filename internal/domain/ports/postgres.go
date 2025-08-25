package ports

import (
	"context"

	"marketflow/internal/domain/models"
)

type PostgresDB interface {
	NewBatchInsert(ctx context.Context, prices []models.PriceStats) error
	GetAverage(context.Context, string, string, string) ([]models.PriceStats, error)
	GetLowHighStat(context.Context, string, string, string, string) (models.PriceStats, error)
	CheckConn() error
}
