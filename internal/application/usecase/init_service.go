package usecase

import (
	"marketflow/internal/core/utils"
	"marketflow/internal/domain/ports"
)

type Service struct {
	Repo  ports.PostgresDB
	Cache ports.Cache
	Valid *utils.Validation
}

func InitService(repo ports.PostgresDB, cache ports.Cache, valid *utils.Validation) ports.ServiceMethods {
	return &Service{
		Repo:  repo,
		Cache: cache,
		Valid: valid,
	}
}
