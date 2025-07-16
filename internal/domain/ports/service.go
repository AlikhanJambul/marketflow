package ports

import (
	"context"
	"marketflow/internal/domain/models"
)

type ServiceMethods interface {
	GetLatestSymService(string) (models.Prices, error)
	GetLatestSymExcService(string, string) (models.Prices, error)
	GetHighestSymService(string) (models.Prices, error)
	GetHighestSymExcService(string, string) (models.Prices, error)
	GetLowestSymService(string) (models.Prices, error)
	GetLowestSymExcService(string, string) (models.Prices, error)
	GetAvgSymService(string) (models.Prices, error)
	GetAvgSymExcService(string, string) (models.Prices, error)
	CheckRedisDb(ctx context.Context) models.HealthResponse
}
