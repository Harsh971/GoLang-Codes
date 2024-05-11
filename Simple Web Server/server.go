package main

import (
	"net/http"
)

func main() {
	// Define a handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	}

	// Register the handler function for the root URL pattern "/"
	http.HandleFunc("/", handler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
