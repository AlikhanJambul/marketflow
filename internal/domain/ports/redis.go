package ports

import (
	"context"
	"marketflow/internal/domain/models"
)

type Cache interface {
	SetLatest(prices models.Prices) error
	GetLatest(string) (models.Prices, error)
	Check(context.Context) error
}
