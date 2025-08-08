package ports

import (
	"context"
	"marketflow/internal/domain/models"
	"time"
)

type PostgresDB interface {
	NewBatchInsert(ctx context.Context, prices []models.PriceStats) error
	GetHighestSym(context.Context, string, time.Duration) (models.PriceStats, error)
	GetHighestSymExc(context.Context, string, string, time.Duration) (models.PriceStats, error)
	GetLowestSym(context.Context, string, time.Duration) (models.PriceStats, error)
	GetLowestSymExc(context.Context, string, string, time.Duration) (models.PriceStats, error)
	GetAvgSym(context.Context, string, time.Duration) (models.PriceStats, error)
	GetAvgSymExc(context.Context, string, string, time.Duration) (models.PriceStats, error)
	CheckConn() error
}
