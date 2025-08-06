package exchange

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"marketflow/internal/domain/models"
	"net"
	"time"
)

type BirgeClient struct {
	name   string
	addr   string
	out    chan<- models.Prices
	conn   net.Conn
	stopCh chan struct{}
}

func NewBirgeClient(name, addr string, out chan<- models.Prices) *BirgeClient {
	return &BirgeClient{
		name:   name,
		addr:   addr,
		out:    out,
		stopCh: make(chan struct{}),
	}
}

func (c *BirgeClient) Start(ctx context.Context) {
	slog.Info("starting birge client", "exchange", c.name, "addr", c.addr)

	for {
		select {
		case <-ctx.Done():
			slog.Info("birge client stopped by context", "exchange", c.name)
			return
		case <-c.stopCh:
			slog.Info("birge client stopped manually", "exchange", c.name)
			return
		default:
			if err := c.connectAndRead(ctx); err != nil {
				slog.Warn("connection failed", "exchange", c.name, "error", err)
				select {
				case <-ctx.Done():
					return
				case <-c.stopCh:
					return
				case <-time.After(5 * time.Second):
					slog.Info("reconnecting...", "exchange", c.name)
				}
			}
		}
	}
}

func (c *BirgeClient) connectAndRead(ctx context.Context) error {
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
			case c.out <- price:
				slog.Debug("sent price", "exchange", c.name, "symbol", data.Symbol, "price", data.Price)
			case <-ctx.Done():
				return ctx.Err()
			case <-c.stopCh:
				return nil
			}
		}
	}
}

func (c *BirgeClient) Stop() {
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
