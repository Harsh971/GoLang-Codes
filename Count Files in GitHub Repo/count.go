package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type RepositoryContents []struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter GitHub repository URL: ")
	repoURL, _ := reader.ReadString('\n')
	repoURL = strings.TrimSpace(repoURL)

	fmt.Print("Enter the file extension (e.g., 'txt', 'go', etc.): ")
	extension, _ := reader.ReadString('\n')
	extension = strings.TrimSpace(extension)

	fileCount, err := getFileCount(repoURL, extension)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Number of %s files in the repository: %d\n", extension, fileCount)
}

func getFileCount(repoURL, extension string) (int, error) {
	// Extract owner and repository name from the URL
	parts := strings.Split(repoURL, "/")
	if len(parts) < 4 {
		return 0, fmt.Errorf("invalid repository URL")
	}
	owner := parts[3]
	repo := parts[4]

	// Fetch repository contents
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/contents", owner, repo))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var contents RepositoryContents
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		return 0, err
	}

	// Count files with the specified extension
	fileCount := 0
	for _, content := range contents {
		if content.Type == "file" && strings.HasSuffix(content.Name, "."+extension) {
			fileCount++
		}
	}

	return fileCount, nil
}
