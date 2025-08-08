package ports

import (
	"context"
	"marketflow/internal/domain/models"
)

type ServiceMethods interface {
	GetLatestSymService(string) (models.LatestPrice, error)
	GetLatestSymExcService(string, string) (models.LatestPrice, error)
	GetHighestSymService(string, string) (models.PriceStats, error)
	GetHighestSymExcService(string, string, string) (models.PriceStats, error)
	GetLowestSymService(string, string) (models.PriceStats, error)
	GetLowestSymExcService(string, string, string) (models.PriceStats, error)
	GetAvgSymService(string, string) (models.PriceStats, error)
	GetAvgSymExcService(string, string, string) (models.PriceStats, error)
	CheckRedisDb(ctx context.Context) models.HealthResponse
}
