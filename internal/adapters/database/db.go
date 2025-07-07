package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"marketflow/internal/core/models"
)

func ConnDb(config models.DB) (*sql.DB, error) {
	postgresConnStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", postgresConnStr)
	if err != nil {
		slog.Error("Не удалось подключится в PostgreSQL")
		return nil, fmt.Errorf("ошибка подключения к PostgreSQL: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		slog.Error("Не удалось подключится в PostgreSQL")
		return nil, fmt.Errorf("не удалось подключиться к PostgreSQL: %w", err)
	}

	slog.Info("Успешное подключение к PostgreSQL!")
	return db, nil
}
