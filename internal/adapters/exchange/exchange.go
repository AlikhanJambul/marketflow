package exchange

import (
	"bufio"
	"fmt"
	"log/slog"
	"marketflow/internal/domain/models"
	"net"
	"strconv"
	"strings"
	"time"
)

func GetDataBirge(sourses []models.Sourse) {
	for _, sours := range sourses {
		go func(addr string, out chan<- models.Prices) {
			defer close(out)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				slog.Warn("birge %s doesn't work", addr)
				return
			}
			defer conn.Close()

			scanner := bufio.NewScanner(conn)

			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)
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

				update := models.Prices{
					Exchange:  addr,
					Value:     value,
					Timestamp: time.Now(),
					PairName:  parts[0],
				}

				out <- update
			}

			if err := scanner.Err(); err != nil {
				slog.Error("scanner error for %s: %v", addr, err)
			}

		}(sours.Addr, sours.SourseChan)
	}

}
