package ports

import (
	"context"
	"marketflow/internal/domain/models"
	"time"
)

type PostgresDB interface {
	NewBatchInsert(ctx context.Context, prices []models.PriceStats) error
	GetAverage(context.Context, string, string, time.Duration) ([]models.PriceStats, error)
	GetLowHighStat(context.Context, string, string, string, time.Duration) (models.PriceStats, error)
	CheckConn() error
}
