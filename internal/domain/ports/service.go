package ports

import (
	"context"
	"marketflow/internal/domain/models"
)

type ServiceMethods interface {
	GetLatestService(string, string) (models.LatestPrice, error)
	GetAvgService(string, string, string) ([]models.PriceStats, error)
	CheckRedisDb(ctx context.Context) models.HealthResponse
	GetStatService(string, string, string, string) (models.PriceStats, error)
}
