package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"marketflow/internal/domain/models"
	"time"

	_ "github.com/lib/pq"
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
		if err != nil {
			slog.Error("Ошибка sql.Open", "err", err)
			time.Sleep(2 * time.Second)
			continue
		}

		if pingErr := db.Ping(); pingErr == nil {
			slog.Info("Успешное подключение к PostgreSQL!")
			return db, nil
		} else {
			slog.Error("Ошибка db.Ping", "err", pingErr)
		}

		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("после 10 попыток не удалось подключиться к PostgreSQL: %w", err)
}
