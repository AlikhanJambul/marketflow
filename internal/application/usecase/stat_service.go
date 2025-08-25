package usecase

import (
	"context"

	"marketflow/internal/core/apperrors"
	"marketflow/internal/core/utils"
	"marketflow/internal/domain/models"
)

func (s *Service) GetStatService(symbol, exchange, price, duration string) (models.PriceStats, error) {
	if ok := s.Valid.CheckAll(symbol, exchange); !ok {
		return models.PriceStats{}, apperrors.ErrInavalidBody
	}

	validDuration, ok := utils.CheckDuration(duration)
	if !ok {
		return models.PriceStats{}, apperrors.ErrDuration
	}

	if err := s.Repo.CheckConn(); err != nil {
		return models.PriceStats{}, apperrors.ErrDB
	}

	result, err := s.Repo.GetLowHighStat(context.Background(), symbol, exchange, price, validDuration)
	if err != nil {
		return models.PriceStats{}, err
	}

	return result, nil
}
