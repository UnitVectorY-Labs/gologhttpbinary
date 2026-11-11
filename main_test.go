package main

import (
	"bytes"
	"encoding/json/v2"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerLogsRequest(t *testing.T) {
	var output bytes.Buffer
	previous := log.Writer()
	log.SetOutput(&output)
	t.Cleanup(func() { log.SetOutput(previous) })

	req := httptest.NewRequest(http.MethodPost, "/example", bytes.NewReader([]byte{0, 255}))
	req.Header.Set("X-Example", "<value>&")
	w := httptest.NewRecorder()
	handler(w, req)

	if w.Code != http.StatusOK || w.Body.String() != "OK\n" || w.Header().Get("X-App-Version") != Version {
		t.Fatalf("unexpected response: status=%d body=%q headers=%v", w.Code, w.Body.String(), w.Header())
	}
	line := output.String()
	start := strings.IndexByte(line, '{')
	if start < 0 {
		t.Fatalf("no JSON log entry: %q", line)
	}
	var entry LogEntry
	if err := json.Unmarshal([]byte(strings.TrimSpace(line[start:])), &entry); err != nil {
		t.Fatalf("invalid JSON log entry: %v", err)
	}
	if entry.BodyBase64 != "AP8=" || entry.Path != "/example" || entry.Headers["X-Example"] != "<value>&" {
		t.Fatalf("unexpected log entry: %+v", entry)
	}
	if !strings.Contains(line, "<value>&") {
		t.Fatalf("expected v2's unescaped HTML characters in log: %q", line)
	}
}

func TestHandlerRejectsInvalidUTF8(t *testing.T) {
	previous := log.Writer()
	log.SetOutput(io.Discard)
	t.Cleanup(func() { log.SetOutput(previous) })

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Example", string([]byte{0xff}))
	w := httptest.NewRecorder()
	handler(w, req)
	if w.Code != http.StatusInternalServerError || w.Body.String() != "Unable to create log entry\n" {
		t.Fatalf("unexpected response: status=%d body=%q", w.Code, w.Body.String())
	}
}
