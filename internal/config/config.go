package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	AutoUpdate          bool   `json:"autoUpdate"`
	UpdateCheckInterval string `json:"updateCheckInterval"` // "daily", "weekly", etc.
	LogLevel            string `json:"logLevel"`            // "debug", "info", "warn"
	CliStyle            string `json:"cliStyle"`            // "plain", "rainbow"
	Telemetry           bool   `json:"telemetry"`
}

var DefaultConfig = Config{
	AutoUpdate:          true,
	UpdateCheckInterval: "daily",
	LogLevel:            "info",
	CliStyle:            "rainbow",
	Telemetry:           false,
}

func LoadConfig() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return DefaultConfig, err
	}
	configPath := filepath.Join(home, ".gobpmn", "config.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		// Datei nicht gefunden? Default verwenden.
		return DefaultConfig, nil
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return DefaultConfig, err
	}

	return cfg, nil
}

func UpdateField(cfg Config, key string, value string) (Config, error) {
	switch key {
	case "autoUpdate":
		v, err := strconv.ParseBool(value)
		if err != nil {
			return cfg, err
		}
		cfg.AutoUpdate = v
	case "updateCheckInterval":
		cfg.UpdateCheckInterval = value
	case "logLevel":
		cfg.LogLevel = value
	case "cliStyle":
		cfg.CliStyle = value
	case "telemetry":
		v, err := strconv.ParseBool(value)
		if err != nil {
			return cfg, err
		}
		cfg.Telemetry = v
	default:
		return cfg, errors.New("unknown key: " + key)
	}
	return cfg, nil
}

func SaveConfig(cfg Config) error {
	home, _ := os.UserHomeDir()
	configPath := filepath.Join(home, ".gobpmn", "config.json")
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(configPath, data, 0644)
}
