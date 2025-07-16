package ports

import (
	"context"
	"marketflow/internal/domain/models"
)

type Cache interface {
	Get(string) (models.Prices, error)
	Set(string, models.Prices) error
	Check(context.Context) error
}
