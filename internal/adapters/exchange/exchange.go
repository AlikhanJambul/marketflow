package exchange

import (
	"bufio"
	"encoding/json"
	"log/slog"
	"marketflow/internal/domain/models"
	"net"
	"strings"
	"time"
)

func GetDataBirge(sources []models.Sourse) {
	for _, source := range sources {
		exchange := strings.Split(source.Addr, ":")
		slog.Info(exchange[0])
		go func(addr string, out chan<- models.Prices) {
			//defer close(out)

			conn, err := net.Dial("tcp", addr)
			if err != nil {
				slog.Warn("birge %s doesn't work", addr)
				return
			}
			defer conn.Close()

			scanner := bufio.NewScanner(conn)

			for scanner.Scan() {
				line := scanner.Text()

				var data struct {
					Symbol    string  `json:"symbol"`
					Price     float64 `json:"price"`
					Timestamp int64   `json:"timestamp"`
				}

				if err := json.Unmarshal([]byte(line), &data); err != nil {
					slog.Warn("Failed to parse JSON from %s: %v", addr, err)
					continue
				}

				out <- models.Prices{
					Exchange:  exchange[0],
					Symbol:    data.Symbol,
					Price:     data.Price,
					Timestamp: time.UnixMilli(data.Timestamp),
				}
			}

			if err := scanner.Err(); err != nil {
				slog.Error("scanner error for %s: %v", addr, err)
			}
		}(source.Addr, source.SourseChan)
	}
}
