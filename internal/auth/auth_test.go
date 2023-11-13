package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
    t.Run("returns error when no Authorization header", func(t *testing.T) {
        headers := http.Header{}
        _, err := GetAPIKey(headers)
        if err == nil {
            t.Errorf("expected an error, got nil")
        }
    })

    t.Run("returns error when Authorization header is malformed", func(t *testing.T) {
        headers := http.Header{}
        headers.Add("Authorization", "malformed")
        _, err := GetAPIKey(headers)
        if err == nil {
            t.Errorf("expected an error, got nil")
        }
    })

    t.Run("returns API key when Authorization header is well formed", func(t *testing.T) {
        headers := http.Header{}
        headers.Add("Authorization", "ApiKey 12345")
        apiKey, err := GetAPIKey(headers)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }
        if apiKey != "12345" {
            t.Errorf("expected API key to be '12345', got '%v'", apiKey)
        }
    })
}