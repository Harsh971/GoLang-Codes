package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// CountryResponse struct to hold information about the response from the API
type CountryResponse []struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Capital []string `json:"capital"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of the country: ")
	countryName, _ := reader.ReadString('\n')
	countryName = strings.TrimSpace(countryName)

	capital, err := getCapital(countryName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("The capital of %s is %s\n", countryName, capital)
}

func getCapital(countryName string) (string, error) {
	apiURL := fmt.Sprintf("https://restcountries.com/v3.1/name/%s", countryName)
	response, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var countryResponse CountryResponse
	err = json.NewDecoder(response.Body).Decode(&countryResponse)
	if err != nil {
		return "", err
	}

	if len(countryResponse) == 0 {
		return "", fmt.Errorf("country not found")
	}

	// Assuming there could be multiple capitals for some countries, returning the first one
	return countryResponse[0].Capital[0], nil
}
