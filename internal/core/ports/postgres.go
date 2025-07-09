package ports

import "marketflow/internal/domain/models"

type PostgresDB interface {
	BatchInsert([]models.Prices) error
}
