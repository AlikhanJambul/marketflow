package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"marketflow/internal/core/ports"
	"marketflow/internal/domain/models"
	"strings"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.PostgresDB {
	return &Repository{db: db}
}

func (r Repository) BatchInsert(prices []models.Prices) error {
	query := "INSERT INTO aggregated_prices (pair_name, exchange, timestamp, value) VALUES "
	args := []interface{}{}
	values := []string{}

	for i, p := range prices {
		pos := i * 4
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d)", pos+1, pos+2, pos+3, pos+4))
		args = append(args, p.PairName, p.Exchange, p.Timestamp, p.Value)
	}

	query += strings.Join(values, ",")
	_, err := r.db.ExecContext(context.Background(), query, args...)
	slog.Info("Работает!!!!!")
	return err
}
