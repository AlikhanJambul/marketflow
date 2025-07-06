package config

import (
	"os"
	"strconv"
)

type Config struct {
	DB        DB
	Redis     Redis
	Port      string
	Exchanges []string
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}
type Redis struct {
	Host     string
	Port     string
	Password string
	DB       string
}

func Load() *Config {
	exchanges := []string{}
	for i := 1; i <= 3; i++ {
		key := "EXCHANGE" + strconv.Itoa(i)
		val := os.Getenv(key)
		if val != "" {
			exchanges = append(exchanges, val)
		}
	}

	return &Config{
		DB: DB{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			User:     getEnv("POSTGRES_USER", "postgress"),
			Password: getEnv("POSTGRES_PASSWORD", "postgress"),
			DBName:   getEnv("POSTGRES_DB", "postgres"),
			SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
		},
		Redis: Redis{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "4444"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnv("REDIS_DB", "0"),
		},
		Port:      getEnv("PORT", "8080"),
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
