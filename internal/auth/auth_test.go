package auth

import (
	"HBHBHUB"
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headerValue    string
		expectedAPIKey string
		expectError    bool
	}{
		{
			name:           "Valid API Key",
			headerValue:    "ApiKey valid-key-123",
			expectedAPIKey: "valid-key-123",
			expectError:    false,
		},
		{
			name:           "No Authorization Header",
			headerValue:    "",
			expectedAPIKey: "",
			expectError:    true,
		},
		{
			name:           "Malformed Header - Missing ApiKey",
			headerValue:    "Bearer token-123",
			expectedAPIKey: "",
			expectError:    true,
		},
		{
			name:           "Malformed Header - No Key",
			headerValue:    "ApiKey",
			expectedAPIKey: "",
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.headerValue != "" {
				req.Header.Set("Authorization", tt.headerValue)
			}

			apiKey, err := GetAPIKey(req.Header)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if apiKey != tt.expectedAPIKey {
					t.Errorf("Expected API key %q, got %q", tt.expectedAPIKey, apiKey)
				}
			}
		})
	}
}
