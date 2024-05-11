package main

import (
	"fmt"
	"net/http"
)

// Map to store shortened URLs
var urlMap = make(map[string]string)

// Handler for URL shortening requests
func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
		return
	}

	// Generate a short URL (you can implement your own logic here)
	shortURL := fmt.Sprintf("/%d", len(urlMap)+1)
	urlMap[shortURL] = url

	fmt.Fprintf(w, "Shortened URL: %s%s", r.Host, shortURL)
}

// Handler for redirecting shortened URLs
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path
	if longURL, ok := urlMap[shortURL]; ok {
		http.Redirect(w, r, longURL, http.StatusFound)
		return
	}

	http.NotFound(w, r)
}

func main() {
	// Serve index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "index.html")
			return
		}
		redirectHandler(w, r)
	})

	// URL shortening endpoint
	http.HandleFunc("/shorten", shortenURLHandler)

	// Start the HTTP server
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
