package main

import "testing"

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "get main paragraph",
			input: `<html><body>
						<p>Outside paragraph.</p>
						<main>
							<p>Main paragraph.</p>
						</main>
						</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name:     "get paragraph",
			input:    "<html><body><p>Test paragraph</p></body></html>",
			expected: "Test paragraph",
		},
		{
			name:     "paragraph not found",
			input:    "<html><body><h4>This is not a paragraph</h4></body></html>",
			expected: "",
		},
		{
			name:     "handle invalid paragraph",
			input:    "<<222p>//>Titile<//<h1533>//>",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.input)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}

}
