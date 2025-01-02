package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppConfig struct {
	DBConfig DBConfig
}

type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"db_name"`
}

func LoadConfig() (*AppConfig, error) {
	file, err := os.Open("./configs/config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &DBConfig{}
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return &AppConfig{
		DBConfig: *config,
	}, nil
}
