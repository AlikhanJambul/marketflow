package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"marketflow/internal/domain/models"
	"marketflow/internal/domain/ports"
	"strings"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.PostgresDB {
	return &Repository{db: db}
}

func (r *Repository) BatchInsert(ctx context.Context, prices []models.Prices) error {
	query := "INSERT INTO birge_prices (symbol, price, timestamp, exchange) VALUES "
	args := []interface{}{}
	values := []string{}

	for i, p := range prices {
		pos := i * 4
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d)", pos+1, pos+2, pos+3, pos+4))
		args = append(args, p.Symbol, p.Price, p.Timestamp, p.Exchange)
	}

	query += strings.Join(values, ",")
	_, err := r.db.ExecContext(ctx, query, args...)
	slog.Info("Работает!!!!!")
	return err
}
