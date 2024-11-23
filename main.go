package main

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type LogEntry struct {
	BodyBase64 string            `json:"bodyBase64"`
	Headers    map[string]string `json:"headers"`
	Path       string            `json:"path"` // Add path field
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}

	// Base64 encode the body
	bodyBase64 := base64.StdEncoding.EncodeToString(body)

	// Collect headers into a map
	headers := make(map[string]string)
	for key, values := range r.Header {
		headers[key] = values[0] // Take the first value for each header
	}

	// Create log entry
	logEntry := LogEntry{
		BodyBase64: bodyBase64,
		Headers:    headers,
		Path:       r.URL.Path, // Add path to log entry
	}

	// Serialize log entry to JSON
	logJSON, err := json.Marshal(logEntry)
	if err != nil {
		http.Error(w, "Unable to create log entry", http.StatusInternalServerError)
		return
	}

	// Log the JSON
	log.Println(string(logJSON))

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Request logged\n"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
