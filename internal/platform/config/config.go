package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string            `json:"appName"`
	Server   ServerConfig      `json:"server"`
	Database DatabaseConfig    `json:"database"`
	Logging  LoggingConfig     `json:"logging"`
	Services map[string]string `json:"services"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type DatabaseConfig struct {
	URL string `json:"url"`
}

type LoggingConfig struct {
	ConsoleLevel string `json:"consoleLevel"`
	FileLevel    string `json:"fileLevel"`
	File         string `json:"file"`
}

func Load(path string) (Config, error) {
	var cfg Config
	bytes, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
