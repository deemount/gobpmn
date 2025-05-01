package main

import (
	"fmt"

	"github.com/deemount/gobpmn/internal/update"
)

func runShowChangelog() {
	release, err := update.FetchLatestRelease()
	if err != nil {
		fmt.Println("Error on load the changelogs:", err)
		return
	}

	fmt.Printf("📝 %s – Changelog:\n\n%s\n", release.TagName, release.Body)
}
