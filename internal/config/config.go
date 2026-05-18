package config

import "os"

type Config struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
}

func ConfigLoad() *Config {
	return &Config{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DBNAME"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}
}
