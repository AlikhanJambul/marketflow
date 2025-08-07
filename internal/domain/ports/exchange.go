package ports

import "context"

type Client interface {
	Start(ctx context.Context) error
	Stop()
}
