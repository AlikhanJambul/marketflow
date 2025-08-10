package usecase

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"marketflow/internal/core/apperrors"
	"marketflow/internal/domain/models"
)

func (s *Service) GetLatestService(symbol, exchange string) (models.LatestPrice, error) {
	if ok := s.Valid.CheckAll(symbol, exchange); !ok {
		return models.LatestPrice{}, apperrors.ErrInavalidBody
	}

	var key string

	if exchange == "" {
		key = fmt.Sprintf("latest/%s", symbol)
	} else {
		key = fmt.Sprintf("latest/%s/%s", exchange, symbol)
	}

	ctx := context.Background()

	err := s.Cache.Check(ctx)
	if err != nil {
		return models.LatestPrice{}, err
	}

	res, err := s.Cache.GetLatest(key)
	if err == redis.Nil {
		return models.LatestPrice{}, apperrors.ErrRedisNil
	}

	if err != nil {
		return models.LatestPrice{}, err
	}

	return res, nil
}
