package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"marketflow/internal/core/apperrors"
	"marketflow/internal/domain/models"
)

func (s *Service) GetLatestBySymbolService(symbol string) (models.Prices, error) {
	if ok := s.Valid.CheckSymbol(symbol); !ok {
		return models.Prices{}, apperrors.ErrInvalidSymbol
	}

	key := fmt.Sprintf("latest/symbol/%s", symbol)

	res, err := s.Cache.Get(key)
	if err != nil {
		slog.Error(err.Error())
		slog.Info("I've not took data from cache")
		res, err := s.Repo.GetLastestBySymbol(context.Background(), symbol)
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
