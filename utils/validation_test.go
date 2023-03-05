package utils

import (
	"testing"
)

func Test_FixURLScheme(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty url",
			input:    "",
			expected: "http://",
		},
		{
			name:     "url with http scheme",
			input:    "http://example.com",
			expected: "http://example.com",
		},
		{
			name:     "url with https scheme",
			input:    "https://example.com",
			expected: "https://example.com",
		},
		{
			name:     "url without scheme",
			input:    "example.com",
			expected: "http://example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FixURLScheme(tt.input)
			if got != tt.expected {
				t.Errorf("got %q, expected %q", got, tt.expected)
			}
		})
	}
}

func Test_ValidateURL(t *testing.T) {
	validURLs := []string{
		"http://example.com",
		"https://example.com",
		"http://example.com/path/to/resource",
		"http://example.com/path/to/resource?query=parameter",
	}
	invalidURLs := []string{
		"example.com",
		"http:/example.com",
		"http://examplecom",
		"://example.com",
		"",
	}

	for _, rawURL := range validURLs {
		if err := ValidateURL(rawURL); err != nil {
			t.Errorf("Expected no error for valid URL %s, but got %v", rawURL, err)
		}
	}

	for _, rawURL := range invalidURLs {
		if err := ValidateURL(rawURL); err == nil {
			t.Errorf("Expected error for invalid URL %s, but got none", rawURL)
		}
	}
}
