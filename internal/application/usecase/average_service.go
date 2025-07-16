package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"marketflow/internal/core/apperrors"
	"marketflow/internal/domain/models"
)

func (s *Service) GetAvgSymService(symbol string) (models.Prices, error) {
	if ok := s.Valid.CheckSymbol(symbol); !ok {
		return models.Prices{}, apperrors.ErrInvalidSymbol
	}

	key := fmt.Sprintf("avg/%s", symbol)

	result, err := s.Cache.Get(key)
	if err != nil {
		slog.Error(err.Error())

		resultRepo, err := s.Repo.GetAvgSym(context.Background(), symbol)
		if err != nil {
			return models.Prices{}, err
		}

		err = s.Cache.Set(key, resultRepo)
		if err != nil {
			return models.Prices{}, err
		}

		return resultRepo, nil
	}

	return result, nil
}

func (s *Service) GetAvgSymExcService(symbol, exchange string) (models.Prices, error) {
	if ok := s.Valid.CheckAll(symbol, exchange); !ok {
		return models.Prices{}, apperrors.ErrInavalidBody
	}

	key := fmt.Sprintf("avg/%s/%s", exchange, symbol)

	result, err := s.Cache.Get(key)
	if err != nil {
		res, err := s.Repo.GetAvgSymExc(context.Background(), symbol, exchange)
		if err != nil {
			return models.Prices{}, err
		}

		err = s.Cache.Set(key, res)
		if err != nil {
			return models.Prices{}, err
		}

		return res, nil
	}

	return result, nil
}
