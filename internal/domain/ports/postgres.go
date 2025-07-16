package ports

import (
	"context"
	"marketflow/internal/domain/models"
)

type PostgresDB interface {
	BatchInsert(context.Context, []models.Prices) error
	GetLastestBySymbol(context.Context, string) (models.Prices, error)
}
