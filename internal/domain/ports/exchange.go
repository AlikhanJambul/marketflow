package ports

import (
	"context"

	"marketflow/internal/domain/models"
)

type Client interface {
	Start(ctx context.Context, out chan<- models.Prices) error
	Stop()
}
