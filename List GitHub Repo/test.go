package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a GitHub username as an argument.")
		return
	}
	username := os.Args[1] // Fetch GitHub username from command-line arguments

	// Make a GET request to the GitHub API to fetch the user's repositories
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching repositories:", err)
		return
	}
	defer resp.Body.Close()

	var repos []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	fmt.Printf("Repositories for user %s:\n", username)
	for _, repo := range repos {
		fmt.Printf("- %s: %s\n", repo.Name, repo.Description)
	}
}
