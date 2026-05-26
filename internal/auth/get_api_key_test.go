package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeySuccess(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if key != "my-secret-key" {
		t.Fatalf("expected key to be 'my-secret-key', got %s", key)
	}
}

func TestGetAPIKeyMissing(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}
