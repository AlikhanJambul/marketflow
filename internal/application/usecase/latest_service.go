package usecase

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"marketflow/internal/core/apperrors"
	"marketflow/internal/domain/models"
)

func (s *Service) GetLatestSymService(symbol string) (models.LatestPrice, error) {
	if ok := s.Valid.CheckSymbol(symbol); !ok {
		return models.LatestPrice{}, apperrors.ErrInvalidSymbol
	}

	key := fmt.Sprintf("latest/%s", symbol)

	res, err := s.Cache.GetLatest(key)
	if err == redis.Nil {
		slog.Info("GetLatestSymService: key not found")
		return models.LatestPrice{}, apperrors.ErrRedisNil
	}
	if err != nil {
		slog.Error(err.Error())
		return models.LatestPrice{}, err
	}

	return res, nil
}

func (s *Service) GetLatestSymExcService(symbol, exchange string) (models.LatestPrice, error) {
	if ok := s.Valid.CheckAll(symbol, exchange); !ok {
		return models.LatestPrice{}, apperrors.ErrInavalidBody
	}

	key := fmt.Sprintf("latest/%s/%s", exchange, symbol)

	res, err := s.Cache.GetLatest(key)
	if err == redis.Nil {
		slog.Info("GetLatestSymExcService: key not found")
		return models.LatestPrice{}, apperrors.ErrRedisNil
	}

	if err != nil {
		slog.Error(err.Error())
		return models.LatestPrice{}, err
	}

	return res, nil
}
