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

func (r *Repository) GetLowHighStat(ctx context.Context, symbol, exchange, price string, duration string) (models.PriceStats, error) {
	orderBy := "min_price ASC"
	if price == "max" {
		orderBy = "max_price DESC"
	}

	query := fmt.Sprintf(`
		SELECT pair_name, exchange, max_price, min_price, timestamp 
		FROM birge_prices
		WHERE pair_name = $1
		  AND timestamp >= NOW() - $2::interval
	`)

	args := []interface{}{symbol, duration}

	if exchange != "" {
		query += " AND exchange = $3"
		args = append(args, exchange)
	}

	query += fmt.Sprintf(" ORDER BY %s LIMIT 1", orderBy)

	var result models.PriceStats
	err := r.db.QueryRowContext(ctx, query, args...).
		Scan(&result.Pair, &result.Exchange, &result.Max, &result.Min, &result.Timestamp)
	if err != nil {
		return models.PriceStats{}, err
	}

	return result, nil
}
