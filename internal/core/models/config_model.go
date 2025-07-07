package models

type Config struct {
	DB        DB
	Redis     Redis
	Port      string
	Exchanges []string
}
