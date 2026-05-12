package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:      "absolute URL",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="https://crawler-test.com/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		{
			name:      "relative URL",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/images/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/images/logo.png"},
		},
		{
			name:     "multiple images",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
				<img src="/images/logo.png" alt="Logo">
				<img src="https://cdn.crawler-test.com/banner.png" alt="Banner">
			</body></html>`,
			expected: []string{"https://crawler-test.com/images/logo.png", "https://cdn.crawler-test.com/banner.png"},
		},
		{
			name:      "trim whitespace",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="  /images/logo.png  " alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/images/logo.png"},
		},
		{
			name:      "no src",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img alt="Logo"></body></html>`,
			expected:  nil,
		},
		{
			name:      "invalid image URL",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="::bad url::" alt="Logo"></body></html>`,
			expected:  nil,
		},
		{
			name:      "bad HTML",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html body><img src="/images/logo.png"></html body>`,
			expected:  []string{"https://crawler-test.com/images/logo.png"},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}

			actual, err := getImagesFromHTML(tc.inputBody, baseURL)

			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URLs %v, got URLs %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}
