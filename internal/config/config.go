package config

import "os"

type Config struct {
	User     string
	Password string
	DBName   string
}

func ConfigLoad() *Config {
	return &Config{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}
}
