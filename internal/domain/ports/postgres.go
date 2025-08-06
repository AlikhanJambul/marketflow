package ports

import (
	"context"
	"marketflow/internal/domain/models"
)

type PostgresDB interface {
	NewBatchInsert(ctx context.Context, prices []models.PriceStats) error
	GetHighestSym(context.Context, string) (models.PriceStats, error)
	GetHighestSymExc(context.Context, string, string) (models.PriceStats, error)
	GetLowestSym(context.Context, string) (models.PriceStats, error)
	GetLowestSymExc(context.Context, string, string) (models.PriceStats, error)
	GetAvgSym(context.Context, string) (models.PriceStats, error)
	GetAvgSymExc(context.Context, string, string) (models.PriceStats, error)
	CheckConn() error
}
