package exchange

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net"
	"time"

	"marketflow/internal/domain/models"
	"marketflow/internal/domain/ports"
)

type LiveClient struct {
	name   string
	addr   string
	conn   net.Conn
	stopCh chan struct{}
}

func NewBirgeClient(name, addr string) ports.Client {
	return &LiveClient{
		name:   name,
		addr:   addr,
		stopCh: make(chan struct{}),
	}
}

func (c *LiveClient) Start(ctx context.Context, out chan<- models.Prices) error {
	slog.Info("starting birge client", "exchange", c.name, "addr", c.addr)

	for {
		select {
		case <-ctx.Done():
			slog.Info("birge client stopped by context", "exchange", c.name)
			return nil
		case <-c.stopCh:
			slog.Info("birge client stopped manually", "exchange", c.name)
			return nil
		default:
			if err := c.connectAndRead(ctx, out); err != nil {
				slog.Warn("connection failed", "exchange", c.name, "error", err)
				select {
				case <-ctx.Done():
					return nil
				case <-c.stopCh:
					return nil
				case <-time.After(5 * time.Second):
					slog.Info("reconnecting...", "exchange", c.name)
				}
			}
		}
	}
}

func (c *LiveClient) connectAndRead(ctx context.Context, out chan<- models.Prices) error {
	conn, err := net.DialTimeout("tcp", c.addr, 5*time.Second)
	if err != nil {
		return err
	}
	c.conn = conn
	defer conn.Close()

	slog.Info("connected to exchange", "exchange", c.name)

	scanner := bufio.NewScanner(conn)

	for {
		conn.SetReadDeadline(time.Now().Add(15 * time.Second))

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.stopCh:
			return nil
		default:
			if !scanner.Scan() {
				if err := scanner.Err(); err != nil {
					return err
				}
				return errors.New("connection closed")
			}

			line := scanner.Text()
			var data struct {
				Symbol    string  `json:"symbol"`
				Price     float64 `json:"price"`
				Timestamp int64   `json:"timestamp"`
			}

			if err := json.Unmarshal([]byte(line), &data); err != nil {
				slog.Warn("failed to unmarshal JSON", "exchange", c.name, "data", line, "error", err)
				continue
			}

			price := models.Prices{
				Exchange:  c.name,
				Symbol:    data.Symbol,
				Price:     data.Price,
				Timestamp: time.UnixMilli(data.Timestamp),
			}

			select {
			case out <- price:
				slog.Debug("sent price", "exchange", c.name, "symbol", data.Symbol, "price", data.Price)
			case <-ctx.Done():
				return ctx.Err()
			case <-c.stopCh:
				return nil
			}
		}
	}
}

func (c *LiveClient) Stop() {
	select {
	case <-c.stopCh:
	default:
		close(c.stopCh)
		slog.Info("client stop requested", "exchange", c.name)
	}

	if c.conn != nil {
		_ = c.conn.Close()
	}
}
