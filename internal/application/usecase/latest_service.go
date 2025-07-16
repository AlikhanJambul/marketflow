package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"marketflow/internal/core/apperrors"
	"marketflow/internal/domain/models"
)

func (s *Service) GetLatestSymService(symbol string) (models.Prices, error) {
	if ok := s.Valid.CheckSymbol(symbol); !ok {
		return models.Prices{}, apperrors.ErrInvalidSymbol
	}

	key := fmt.Sprintf("latest/%s", symbol)

	res, err := s.Cache.Get(key)
	if err != nil {
		slog.Error(err.Error())
		slog.Info("I've not took data from cache")

		resRepo, err := s.Repo.GetLastestSym(context.Background(), symbol)
		if err != nil {
			return models.Prices{}, err
		}

		err = s.Cache.Set(key, resRepo)
		if err != nil {
			slog.Error(err.Error())
		}

		return resRepo, nil
	}

	slog.Info("I've took data from cache")
	return res, nil
}

func (s *Service) GetLatestSymExcService(symbol, exchange string) (models.Prices, error) {
	if ok := s.Valid.CheckAll(symbol, exchange); !ok {
		return models.Prices{}, apperrors.ErrInavalidBody
	}

	key := fmt.Sprintf("latest/%s/%s", exchange, symbol)

	res, err := s.Cache.Get(key)
	if err != nil {
		slog.Error(err.Error())
		slog.Info("I've not took data from cache")

		res, err := s.Repo.GetLatestSymExc(context.Background(), symbol, exchange)
		if err != nil {
			return models.Prices{}, err
		}

		err = s.Cache.Set(key, res)
		if err != nil {
			return models.Prices{}, err
		}

		return res, nil
	}

	slog.Info("I've took data from cache")
	return res, nil
}
