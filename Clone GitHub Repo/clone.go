package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Prompt user to input GitHub repository link
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter GitHub repository link: ")
	repoLink, _ := reader.ReadString('\n')

	// Remove leading/trailing whitespace and newlines
	repoLink = strings.TrimSpace(repoLink)

	// Clone the repository
	if err := cloneRepository(repoLink); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Repository cloned successfully!")
}

func cloneRepository(repoLink string) error {
	// Execute git clone command
	cmd := exec.Command("git", "clone", repoLink)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
