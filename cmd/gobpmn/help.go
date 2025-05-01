package main

import "fmt"

func PrintHelp() {
	fmt.Println("Usage: gobpmn <command>")
	fmt.Println("\nCommands available:")
	fmt.Println("  model                     → start the bpmn modeler")
	fmt.Println("  test                      → perform tests")
	fmt.Println("  debug                     → start debug mode")
	fmt.Println("  help                      → display this help")
	fmt.Println("  changelog                 → show changelog")
	fmt.Println("  release                   → show latest release information")
	fmt.Println("  config                    → show or modify the configuration")
	fmt.Println("  version                   → show the application version")
	fmt.Println("  update                    → update the application")
	fmt.Println("  update -check             → check for updates")
	fmt.Println("  update -version <version> → update to a specific version")
}
