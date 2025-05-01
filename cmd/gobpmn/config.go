package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/deemount/gobpmn/internal/config"
)

func RunConfigCommand(args []string) {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	if len(args) == 0 {
		fmt.Println("Current configuration:")
		printConfig(cfg)
		return
	}

	if len(args) == 3 && args[0] == "set" {
		key := args[1]
		value := args[2]

		updated, err := config.UpdateField(cfg, key, value)
		if err != nil {
			fmt.Println("Error changing configuration:", err)
			os.Exit(1)
		}

		err = config.SaveConfig(updated)
		if err != nil {
			fmt.Println("Error saving configuration:", err)
			os.Exit(1)
		}

		fmt.Println("Configuration updated.")
		return
	}

	fmt.Println("  Invalid arguments. Utilization:")
	fmt.Println("  gobpmn config               # shows configuration")
	fmt.Println("  gobpmn config set <key> <value>")
}

func EnsureConfigExists() error {
	home, _ := os.UserHomeDir()
	configDir := filepath.Join(home, ".gobpmn")
	configPath := filepath.Join(configDir, "config.json")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
		data, _ := json.MarshalIndent(config.DefaultConfig, "", "  ")
		return os.WriteFile(configPath, data, 0644)
	}
	return nil
}

func printConfig(cfg config.Config) {
	fmt.Println("  autoUpdate:", cfg.AutoUpdate)
	fmt.Println("  updateCheckInterval:", cfg.UpdateCheckInterval)
	fmt.Println("  logLevel:", cfg.LogLevel)
	fmt.Println("  cliStyle:", cfg.CliStyle)
	fmt.Println("  telemetry:", cfg.Telemetry)
}
