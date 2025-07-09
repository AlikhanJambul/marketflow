package config

import (
	"marketflow/internal/domain/models"
	"os"
	"strconv"
)

func Load() *models.Config {
	exchanges := []string{}
	for i := 1; i <= 3; i++ {
		key := "EXCHANGE" + strconv.Itoa(i)
		val := os.Getenv(key)
		if val != "" {
			exchanges = append(exchanges, val)
		}
	}

	return &models.Config{
		DB: models.DB{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", "postgres"),
			DBName:   getEnv("POSTGRES_DB", "postgres"),
			SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
		},
		Redis: models.Redis{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "4444"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvInt("REDIS_DB", 0),
		},
		Port:      getEnv("SERVER_PORT", "8080"),
		Exchanges: exchanges,
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func getEnvInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return i
}
