package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("error trying to fetch %v", rawURL)
	}

	req.Header.Set("User-Agent", "BootCrawler/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("got network error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return "", fmt.Errorf("got HTTP error: %s", resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non-HTML response: %s", contentType)
	}

	htmlBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}

	htmlBody := string(htmlBodyBytes)
	return htmlBody, nil
}
