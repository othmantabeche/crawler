package main

import "testing"

func TestConfigureValidatesLimits(t *testing.T) {
	tests := []struct {
		name            string
		maxConcurrency  int
		maxPages        int
		expectedErrText string
	}{
		{name: "reject zero concurrency", maxConcurrency: 0, maxPages: 1, expectedErrText: "maxConcurrency must be at least 1"},
		{name: "reject zero pages", maxConcurrency: 1, maxPages: 0, expectedErrText: "maxPages must be at least 1"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := configure("https://crawler-test.com", tc.maxConcurrency, tc.maxPages)
			if err == nil || err.Error() != tc.expectedErrText {
				t.Fatalf("expected %q, got %v", tc.expectedErrText, err)
			}
		})
	}
}
