package main

import (
	"fmt"

	"github.com/deemount/gobpmn/internal/update"
)

func runShowReleaseInfo() {
	release, err := update.FetchLatestRelease()
	if err != nil {
		fmt.Println("Could not load release:", err)
		return
	}

	fmt.Println("Latest Release:")
	fmt.Println("  Version:   ", release.TagName)
	fmt.Println("  Name:      ", release.Name)
	fmt.Println("  Published: ", release.PublishedAt.Format("2006-01-02"))
	fmt.Println()
	fmt.Println("📝 Changelog:\n", release.Body)
}
