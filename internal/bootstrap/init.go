package bootstrap

import (
	"log/slog"
	"marketflow/internal/adapters/handlers"
	"marketflow/internal/adapters/postgres"
	"marketflow/internal/adapters/redis"
	"marketflow/internal/application/usecase"
	"marketflow/internal/core/config"
	"marketflow/internal/core/utils"
	"marketflow/internal/domain/models"
	"marketflow/internal/domain/ports"
	"net/http"
	"os"
)

var (
	Cfg      *models.Config
	Repo     ports.PostgresDB
	Cache    ports.Cache
	Mux      *http.ServeMux
	Valid    *utils.Validation
	Services ports.ServiceMethods
)

func init() {
	Cfg = config.Load()

	db, err := postgres.ConnDb(Cfg.DB)
	if err != nil {
		slog.Error("PostgreSQL не подключён:", err)
		os.Exit(1)
	}
	Repo = postgres.NewRepository(db)

	rdb := redis.ConnRedis(Cfg.Redis)
	Cache = redis.NewRedisCache(rdb)
	Valid = utils.NewValidation()

	Mux = handlers.InitNewServer()

	Services = usecase.InitService(Repo, Cache, Valid)

	slog.Info("Инициализация завершена успешно")
}
