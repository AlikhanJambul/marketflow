package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"marketflow/internal/core/apperrors"
	"marketflow/internal/core/utils"
	"marketflow/internal/domain/models"
)

func (s *Service) GetHighestSymService(symbol, duration string) (models.PriceStats, error) {
	if ok := s.Valid.CheckSymbol(symbol); !ok {
		return models.PriceStats{}, apperrors.ErrInvalidSymbol
	}

	validDuration, ok := utils.CheckDuration(duration)
	if !ok {
		return models.PriceStats{}, errors.New("invalid duration")
	}

	key := fmt.Sprintf("highest/%s", symbol)

	result, err := s.Cache.Get(key)
	if err != nil {
		slog.Error(err.Error())

		resultRepo, err := s.Repo.GetHighestSym(context.Background(), symbol, validDuration)
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

func (s *Service) GetHighestSymExcService(symbol, exchange, duration string) (models.PriceStats, error) {
	if ok := s.Valid.CheckAll(symbol, exchange); !ok {
		return models.PriceStats{}, apperrors.ErrInavalidBody
	}

	validDuration, ok := utils.CheckDuration(duration)
	if !ok {
		return models.PriceStats{}, errors.New("invalid duration")
	}

	key := fmt.Sprintf("highest/%s/%s", exchange, symbol)

	result, err := s.Cache.Get(key)
	if err != nil {
		res, err := s.Repo.GetHighestSymExc(context.Background(), symbol, exchange, validDuration)
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
