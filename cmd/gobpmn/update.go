package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/deemount/gobpmn/internal/update"
)

func handleUpdate(args []string) {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	manualVersion := updateCmd.String("version", "", "Update to certain version")
	check := updateCmd.Bool("check", false, "Check if new version is available")
	updateCmd.Parse(args)

	if *check {
		latest, newer, err := update.CheckForUpdate(version)
		if err != nil {
			fmt.Println("Error during check:", err)
			return
		}
		if newer {
			fmt.Printf("New version available: %s → %s\n", version, latest)
		} else {
			fmt.Println("You already have the latest version:", version)
		}
		return
	}

	if err := update.RunUpdate(*manualVersion); err != nil {
		fmt.Println("Error during update:", err)
	} else {
		fmt.Println("gobpmn successfully updated!")
	}
}

func runAutoUpdate(force bool) {
	if strings.Contains(version, "dev") {
		return // no update check in dev versions
	}
	// check path (e. g. ~/.gobpmn/.last_update_check)
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	checkFile := filepath.Join(home, ".gobpmn", ".last_update_check")

	if !force && !shouldCheckToday(checkFile) {
		return
	}
	latest, newer, err := update.CheckForUpdate(version)
	if err != nil || !newer {
		if force {
			fmt.Println("You are using the latest version")
		}
		return
	}

	if newer {
		if force {
			fmt.Println("Start automatic update ...")
			if err := update.PerformUpdate(); err != nil {
				fmt.Println("Update failed:", err)
			} else {
				fmt.Println("Update successfully completed!")
			}
		} else {
			fmt.Println("Execute `gobpmn update` to update")
		}
	}

	fmt.Printf(
		"\nA new version (%s) is available! You are using: %s\nRun `gobpmn update` to update.\n\n",
		latest, version,
	)

	_ = updateCheckTimestamp(checkFile)

}

func shouldCheckToday(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return true // no entry available, so check
	}

	lastCheck, err := time.Parse(time.RFC3339, string(data))
	if err != nil {
		return true
	}

	now := time.Now()
	return now.Sub(lastCheck) > 24*time.Hour
}

func updateCheckTimestamp(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return os.WriteFile(path, []byte(time.Now().Format(time.RFC3339)), 0644)
}
