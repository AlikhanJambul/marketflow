package mode

import (
	"context"
	"errors"
	"log/slog"
	"marketflow/internal/adapters/exchange"
	"marketflow/internal/domain/models"
	"marketflow/internal/domain/ports"
	"strings"
	"sync"
)

type Mode string

const (
	Live Mode = "live"
	Test Mode = "test"
)

type Manager struct {
	Cfg     *models.Config
	Mu      sync.Mutex
	Mode    Mode
	Clients []ports.Client
}

func NewManager(cfg *models.Config, mode Mode) *Manager {
	return &Manager{
		Cfg:  cfg,
		Mode: mode,
	}
}

func (m *Manager) Start(ctx context.Context, mode Mode) ([]<-chan models.Prices, error) {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	if m.Mode == mode {
		slog.Info("Mode already set to %s", string(mode))
		return nil, errors.New("already set")
	}

	for _, client := range m.Clients {
		client.Stop()
	}
	m.Clients = nil

	var clients []ports.Client
	chans := []<-chan models.Prices{}

	switch mode {
	case Live:
		var sourceArr []models.Sourse

		for _, addr := range m.Cfg.Exchanges {
			ch := make(chan models.Prices, 15)
			sourceArr = append(sourceArr, models.Sourse{
				SourseChan: ch,
				Addr:       addr,
			})
		}

		for _, source := range sourceArr {
			chans = append(chans, source.SourseChan)
		}

		for _, s := range sourceArr {
			parts := strings.Split(s.Addr, ":")
			exchangeCount := parts[0]

			client := exchange.NewBirgeClient(exchangeCount, s.Addr, s.SourseChan)
			clients = append(clients, client)
		}

	case Test:
		for _, addr := range m.Cfg.Exchanges {
			out := make(chan models.Prices, 15)
			parts := strings.Split(addr, ":")
			exchangeCount := parts[0]

			chans = append(chans, out)

			client := exchange.NewTestClient(out, exchangeCount)
			clients = append(clients, client)
		}

	}

	m.Clients = clients

	for _, client := range clients {
		go func(client ports.Client) {
			if err := client.Start(ctx); err != nil {
				slog.Error(err.Error())
				return
			}
		}(client)
	}

	m.Mode = mode

	slog.Info("mode set to %s", string(m.Mode))
	return chans, nil
}
