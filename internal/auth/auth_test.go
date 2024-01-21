package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test case 1: Valid Authorization header
	headers := http.Header{"Authorization": []string{"ApiKey my-api-ke"}}
	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if apiKey != "my-api-key" {
		t.Errorf("Expected API key 'my-api-key', got '%s'", apiKey)
	}

	// Test case 2: Missing Authorization header
	headers = http.Header{}
	_, err = GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected error '%v', got %v", ErrNoAuthHeaderIncluded, err)
	}

	// Test case 3: Malformed Authorization header
	headers = http.Header{"Authorization": []string{"Bearer token"}}
	_, err = GetAPIKey(headers)
	if err.Error() != "malformed authorization header" {
		t.Errorf("Expected error 'malformed authorization header', got %v", err)
	}
}
