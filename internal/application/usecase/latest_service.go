package usecase

import (
	"context"
	"fmt"
	"marketflow/internal/core/apperrors"
	"marketflow/internal/domain/models"

	"github.com/redis/go-redis/v9"
)

func (s *Service) GetLatestService(symbol, exchange string) (models.Prices, error) {
	if ok := s.Valid.CheckAll(symbol, exchange); !ok {
		return models.Prices{}, apperrors.ErrInavalidBody
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
		return models.Prices{}, apperrors.ErrRedis
	}

	res, err := s.Cache.GetLatest(key)
	if err == redis.Nil {
		return models.Prices{}, apperrors.ErrRedisNil
	}

	if err != nil {
		return models.Prices{}, err
	}

	return res, nil
}
