package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	DBConfig DBConfig
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

func LoadConfig() (*AppConfig, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		DBConfig: DBConfig{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			DBName:   os.Getenv("DB_NAME"),
		},
	}, nil
}
