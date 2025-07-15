package ports

import "marketflow/internal/domain/models"

type ServiceMethods interface {
	GetLatestBySymbolService(symbol string) (models.Prices, error)
}
