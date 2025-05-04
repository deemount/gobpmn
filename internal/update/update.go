package update

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const GitHubRepo = "deemount/gobpmn" // GitHub repository for the project
const GitHubAPIURL = "https://api.github.com/repos/%s/releases/latest"

// GitHubRelease represents the structure of a GitHub release
type GitHubRelease struct {
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	Body        string    `json:"body"`
	PublishedAt time.Time `json:"published_at"`
}

// GetModulePath reads the module path from go.mod
func GetModulePath() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module")), nil
		}
	}
	return "", nil
}

// CheckForUpdate compares current and last version
func CheckForUpdate(currentVersion string) (string, bool, error) {
	url := fmt.Sprintf(GitHubAPIURL, GitHubRepo)
	// Make a GET request to the GitHub API
	// to get the latest release information
	// and check for errors
	resp, err := http.Get(url)
	if err != nil {
		return "", false, err
	}
	defer resp.Body.Close()

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", false, err
	}

	latest := strings.TrimPrefix(release.TagName, "v")
	current := strings.TrimPrefix(currentVersion, "v")

	return release.TagName, latest != current, nil
}

// RunUpdate executes the update - optionally with a specific version
func RunUpdate(version string) error {
	modulePath, err := GetModulePath()
	if err != nil {
		return fmt.Errorf("module path not found: %v", err)
	}

	if version == "" {
		version = "latest"
	}

	fmt.Printf("Update %s@%s...\n", modulePath, version)

	cmd := exec.Command("go", "install", modulePath+"@"+version)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func PerformUpdate() error {
	cmd := exec.Command("sh", "-c", "go install $(grep module go.mod | cut -d ' ' -f 2)@latest")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func FetchLatestRelease() (*GitHubRelease, error) {
	return fetchLatestRelease(GitHubRepo)
}

func fetchLatestRelease(repo string) (*GitHubRelease, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)

	req, _ := http.NewRequest("GET", url, nil)

	// use gitHub token, if set
	// this is needed for private repositories
	// or if the rate limit is reached
	// see https://docs.github.com/en/rest/overview/resources-in-the-rest-api#current-version
	// and https://docs.github.com/en/rest/overview/resources-in-the-rest-api#rate-limiting
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http-error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gitHub api returned: %s", resp.Status)
	}

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("parsing failed with error: %w", err)
	}

	return &release, nil
}
