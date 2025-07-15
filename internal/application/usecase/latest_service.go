package usecase

import (
	"context"
	"errors"
	"marketflow/internal/domain/models"
)

func (s *Service) GetLatestBySymbolService(symbol string) (models.Prices, error) {
	if ok := s.Valid.CheckSymbol(symbol); !ok {
		return models.Prices{}, errors.New("invalid symbol")
	}

	res, err := s.Cache.Get("lastest/symbol")
	if err != nil {
		res, err := s.Repo.GetLastestBySymbol(context.Background())
		if err != nil {
			return models.Prices{}, err
		}

		err = s.Cache.Set("latest/symbol", res)
		if err != nil {
			return models.Prices{}, err
		}

		return res, nil
	}

	return res, nil
}
