package usecase

import (
	"context"
	"marketflow/internal/domain/models"
)

func (s *Service) CheckRedisDb(ctx context.Context) models.HealthResponse {
	status, pgStatus, redisStatus := "", "", ""

	if err := s.Cache.Check(ctx); err != nil {
		redisStatus = "down"
	} else {
		redisStatus = "up"
	}

	if err := s.Repo.CheckConn(); err != nil {
		pgStatus = "down"
	} else {
		pgStatus = "up"
	}

	if pgStatus == "up" && redisStatus == "up" {
		status = "up"
	} else {
		status = "down"
	}

	res := models.HealthResponse{
		Status:   status,
		Redis:    redisStatus,
		Postgres: pgStatus,
	}

	return res
}
