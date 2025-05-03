package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/deemount/gobpmn/internal/banner"
	"github.com/deemount/gobpmn/internal/config"
)

var version = "v0.2.0-dev" // overwritten during build

var AppConfig config.Config

func init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Config could not be loaded. Default is used.")
	}
	AppConfig = cfg
}

func main() {

	log.Printf("gobpmn %s starting ...\n", version)
	if err := EnsureConfigExists(); err != nil {
		log.Fatalf("Error creating config directory: %v", err)
	}
	args := os.Args[1:]

	// force manual check
	forceUpdate := slices.Contains(args, "--force-update-check")

	if AppConfig.AutoUpdate {
		runAutoUpdate(forceUpdate)
	}

	if AppConfig.CliStyle == "rainbow" {
		banner.PrintRainbowBanner()
	}

	if len(os.Args) < 2 {
		PrintHelp()
		return
	}

	cmd := os.Args[1]

	if len(args) > 0 {
		switch cmd {
		case "help":

			PrintHelp()
			return

		case "update":

			handleUpdate(args[1:])
			return

		case "release":

			runShowReleaseInfo()
			return

		case "config":

			configCmd := flag.NewFlagSet("config", flag.ExitOnError)
			if err := configCmd.Parse(os.Args[2:]); err != nil {
				log.Fatalf("Error parsing config command: %v", err)
			}
			return

		case "changelog":

			runShowChangelog()
			return

		case "--version", "-v", "version":

			fmt.Printf("gobpmn version: %s\n", version)
			return

		case "model":

			modelerCmd := flag.NewFlagSet("model", flag.ExitOnError)
			if err := modelerCmd.Parse(os.Args[2:]); err != nil {
				log.Fatalf("Error parsing model command: %v", err)
			}
			//modeler.Start()
			return

		case "test":

			testCmd := flag.NewFlagSet("test", flag.ExitOnError)
			if err := testCmd.Parse(os.Args[2:]); err != nil {
				log.Fatalf("Error parsing test command: %v", err)
			}
			//runner.RunTests()
			return

		case "debug":

			debugCmd := flag.NewFlagSet("debug", flag.ExitOnError)
			if err := debugCmd.Parse(os.Args[2:]); err != nil {
				log.Fatalf("Error parsing debug command: %v", err)
			}
			//debug.Start()
			return

		default:

			fmt.Printf("Unknown command: %s\n", cmd)
			PrintHelp()

		}
	}

	ShowMainMenu()

}

func ShowMainMenu() {

	for {
		fmt.Println(`
			[1] Start modeler
			[2] Execute tests
			[3] Debug mode
			[4] Show help
			[5] Check manual update
			[6] Show changelog
			[7] Show config
			[8] Show version
			[9] Show release info
			[0] Exit
		`)

		fmt.Print("Please select: ")
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			// run modeler
		case "2":
			// run tests
		case "3":
			// run debug
		case "4":
			// show help
		case "5":
			// runManualUpdateCheck
		case "6":
			// show changelog
		case "7":
			// show config
		case "8":
			fmt.Printf("gobpmn version: %s\n", version)
		case "9":
			// show release info
		case "0":
			fmt.Println("See you soon!")
			os.Exit(0)
		default:
			fmt.Println("Invalid input")
		}
	}
}
