package mode

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"

	"marketflow/internal/adapters/exchange"
	"marketflow/internal/domain/models"
	"marketflow/internal/domain/ports"
)

//type Mode string
//
//const (
//	Live Mode = "live"
//	Test Mode = "test"
//)

type Manager struct {
	Cfg     *models.Config
	Mu      sync.Mutex
	Mode    string
	Clients []ports.Client
}

func NewManager(cfg *models.Config, mode string) *Manager {
	return &Manager{
		Cfg:  cfg,
		Mode: mode,
	}
}

func (m *Manager) Start(ctx context.Context, mode string, out chan<- models.Prices) error {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	if m.Mode == mode {
		slog.Info("Mode already set to %s", string(mode))
		return errors.New("already set")
	}

	for _, client := range m.Clients {
		if client == nil {
			slog.Error("Client not initialized")
		}
		client.Stop()
	}
	m.Clients = nil

	var clients []ports.Client

	switch mode {
	case "live":
		for _, addr := range m.Cfg.Exchanges {
			fmt.Println(addr)
			parts := strings.Split(addr, ":")
			exchangeCount := parts[0]

			client := exchange.NewBirgeClient(exchangeCount, addr)
			clients = append(clients, client)
		}

	case "test":
		for _, addr := range m.Cfg.Exchanges {
			fmt.Println(addr)
			parts := strings.Split(addr, ":")
			exchangeCount := parts[0]

			client := exchange.NewTestClient(exchangeCount)
			fmt.Println(client)
			clients = append(clients, client)
		}
	default:
		slog.Info("wrong mode")
		return errors.New("wrong mode")
	}

	m.Clients = clients

	for _, client := range clients {
		go func(client ports.Client) {
			if err := client.Start(ctx, out); err != nil {
				slog.Error(err.Error())
				return
			}
		}(client)
	}

	fmt.Println(m.Clients)

	m.Mode = mode

	slog.Info("mode set to %s", string(m.Mode))
	return nil
}
