package bootstrap

import (
	"log/slog"
	"marketflow/internal/adapters/handlers"
	"marketflow/internal/adapters/postgres"
	"marketflow/internal/adapters/redis"
	"marketflow/internal/application/mode"
	"marketflow/internal/application/usecase"
	"marketflow/internal/core/config"
	"marketflow/internal/core/utils"
	"marketflow/internal/domain/models"
	"marketflow/internal/domain/ports"
	"os"
)

var (
	Cfg   *models.Config
	Repo  ports.PostgresDB
	Cache ports.Cache
	//Mux      *http.ServeMux
	Valid    *utils.Validation
	Services ports.ServiceMethods
	Handlers *handlers.Handler
	Manager  *mode.Manager
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

	Manager = mode.NewManager(Cfg, "test")
	Services = usecase.InitService(Repo, Cache, Valid)
	Handlers = handlers.InitHandlers(Services, Manager)

	slog.Info("Инициализация завершена успешно")
}
