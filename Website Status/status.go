package main

import (
	"fmt"
	"net/http"
)

func checkWebsite(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error checking website: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Website %s is active\n", url)
		return true
	} else {
		fmt.Printf("Website %s is down (status code: %d)\n", url, resp.StatusCode)
		return false
	}
}

func main() {
	var websiteURL string
	fmt.Print("Enter website URL: ")
	fmt.Scanln(&websiteURL)

	isActive := checkWebsite(websiteURL)
	if isActive {
		fmt.Printf("The website %s is active\n", websiteURL)
	} else {
		fmt.Printf("The website %s is down\n", websiteURL)
	}
}
