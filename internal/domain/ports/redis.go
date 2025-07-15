package ports

import "marketflow/internal/domain/models"

type Cache interface {
	Get(string) (models.Prices, error)
	Set(string, models.Prices) error
}
