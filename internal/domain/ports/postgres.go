package ports

import (
	"context"
	"marketflow/internal/domain/models"
)

type PostgresDB interface {
	BatchInsert(context.Context, []models.Prices) error
	GetLastestSym(context.Context, string) (models.Prices, error)
	GetLatestSymExc(context.Context, string, string) (models.Prices, error)
	GetHighestSym(context.Context, string) (models.Prices, error)
	GetHighestSymExc(context.Context, string, string) (models.Prices, error)
	GetLowestSym(context.Context, string) (models.Prices, error)
	GetLowestSymExc(context.Context, string, string) (models.Prices, error)
	GetAvgSym(context.Context, string) (models.Prices, error)
	GetAvgSymExc(context.Context, string, string) (models.Prices, error)
	CheckConn() error
}
