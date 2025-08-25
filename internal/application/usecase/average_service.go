package usecase

import (
	"context"

	"marketflow/internal/core/apperrors"
	"marketflow/internal/core/utils"
	"marketflow/internal/domain/models"
)

func (s *Service) GetAvgService(symbol, exchange, duration string) ([]models.PriceStats, error) {
	if ok := s.Valid.CheckAll(symbol, exchange); !ok {
		return nil, apperrors.ErrInavalidBody
	}

	validDuration, ok := utils.CheckDuration(duration)
	if !ok {
		return nil, apperrors.ErrDuration
	}

	if err := s.Repo.CheckConn(); err != nil {
		return nil, apperrors.ErrDB
	}

	result, err := s.Repo.GetAverage(context.Background(), symbol, exchange, validDuration)
	if err != nil {
		return nil, err
	}

	return result, nil
}
