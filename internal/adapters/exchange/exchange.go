package exchange

import (
	"bufio"
	"log/slog"
	"marketflow/internal/core/models"
	"net"
	"strconv"
	"strings"
	"time"
)

func GetDataBirge(birge []string, out chan<- models.Exchange) {
	for _, addr := range birge {
		go func(addr string) {
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				slog.Warn("birge %s doesn't work", addr)
				return
			}
			defer conn.Close()

			scanner := bufio.NewScanner(conn)

			for scanner.Scan() {
				line := scanner.Text()

				parts := strings.Fields(line)

				if len(parts) != 2 {
					slog.Warn("Invalid data")
					continue
				}

				value, err := strconv.ParseFloat(parts[1], 64)
				if err != nil {
					slog.Warn("Invalid value")
					continue
				}

				update := models.Exchange{
					Source:    addr,
					Value:     value,
					Timestamp: time.Now(),
					Symbol:    parts[0],
				}

				out <- update
			}

			if err := scanner.Err(); err != nil {
				slog.Error("scanner error for %s: %v", addr, err)
			}

		}(addr)
	}

}
