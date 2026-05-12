package main

import (
	"testing"
)

func TestGetHeadingFromHTMLBasic(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "get h1 title",
			input:    "<html><body><h1>Test Title H1</h1></body></html>",
			expected: "Test Title H1",
		},
		{
			name:     "get h2 title",
			input:    "<html><body><h2>Test Title H2</h2></body></html>",
			expected: "Test Title H2",
		},
		{
			name:     "titile not found",
			input:    "<html><body><p>This is not a title</p></body></html>",
			expected: "",
		},
		{
			name:     "handle invalid HTML",
			input:    "<<h12>//>Titile<//<h1533>//>",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.input)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
