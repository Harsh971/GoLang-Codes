package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Function to check if the email format is valid
func isValidEmailFormat(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	local := parts[0]
	domain := parts[1]

	if len(local) == 0 || len(domain) == 0 {
		return false
	}

	if strings.Contains(local, " ") || strings.Contains(domain, " ") {
		return false
	}

	domainParts := strings.Split(domain, ".")
	if len(domainParts) < 2 {
		return false
	}

	for _, part := range domainParts {
		if len(part) == 0 {
			return false
		}
	}

	return true
}

// Function to check if the domain of the email exists
func isDomainValid(domain string) bool {
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return false
	}
	return true
}

// Function to verify the email address
func verifyEmail(email string) bool {
	if !isValidEmailFormat(email) {
		return false
	}

	parts := strings.Split(email, "@")
	domain := parts[1]

	return isDomainValid(domain)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter an email address: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Remove trailing newline character
	email = strings.TrimSpace(email)

	if verifyEmail(email) {
		fmt.Printf("Email %s is valid.\n", email)
	} else {
		fmt.Printf("Email %s is not valid or domain does not exist.\n", email)
	}
}
