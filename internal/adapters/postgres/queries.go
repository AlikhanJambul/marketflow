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

func (r *Repository) NewBatchInsert(ctx context.Context, prices []models.PriceStats) error {
	query := "INSERT INTO birge_prices (exchange, pair_name, timestamp, average_price, min_price, max_price) VALUES "
	args := []interface{}{}
	values := []string{}

	for i, p := range prices {
		pos := i * 6
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", pos+1, pos+2, pos+3, pos+4, pos+5, pos+6))
		args = append(args, p.Exchange, p.Pair, p.Timestamp, p.Average, p.Min, p.Max)
	}

	query += strings.Join(values, ",")

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		slog.Error("error:", err)
	}

	return err
}

func (r *Repository) CheckConn() error {
	return r.db.Ping()
}
