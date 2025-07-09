package models

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
	DB       int
}
