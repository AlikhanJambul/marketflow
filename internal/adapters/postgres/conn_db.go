package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"marketflow/internal/domain/models"
	"time"
)

func ConnDb(config models.DB) (*sql.DB, error) {
	postgresConnStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName,
	)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", postgresConnStr)
		if err == nil && db.Ping() == nil {
			slog.Info("Успешное подключение к PostgreSQL!")
			return db, nil
		}
		slog.Error(err.Error())
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("после 10 попыток не удалось подключиться к PostgreSQL: %w", err)
}
