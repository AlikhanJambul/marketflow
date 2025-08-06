package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"marketflow/internal/core/apperrors"
	"marketflow/internal/domain/models"
)

func (s *Service) GetHighestSymService(symbol string) (models.PriceStats, error) {
	if ok := s.Valid.CheckSymbol(symbol); !ok {
		return models.PriceStats{}, apperrors.ErrInvalidSymbol
	}

	key := fmt.Sprintf("highest/%s", symbol)

	result, err := s.Cache.Get(key)
	if err != nil {
		slog.Error(err.Error())

		resultRepo, err := s.Repo.GetHighestSym(context.Background(), symbol)
		if err != nil {
			return models.PriceStats{}, err
		}

		err = s.Cache.Set(key, resultRepo)
		if err != nil {
			return models.PriceStats{}, err
		}

		return resultRepo, nil
	}

	return result, nil
}

func (s *Service) GetHighestSymExcService(symbol, exchange string) (models.PriceStats, error) {
	if ok := s.Valid.CheckAll(symbol, exchange); !ok {
		return models.PriceStats{}, apperrors.ErrInavalidBody
	}

	key := fmt.Sprintf("highest/%s/%s", exchange, symbol)

	result, err := s.Cache.Get(key)
	if err != nil {
		res, err := s.Repo.GetHighestSymExc(context.Background(), symbol, exchange)
		if err != nil {
			return models.PriceStats{}, err
		}

		err = s.Cache.Set(key, res)
		if err != nil {
			return models.PriceStats{}, err
		}

		return res, nil
	}

	return result, nil
}
