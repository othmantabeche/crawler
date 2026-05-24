package main

import (
	"errors"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", errors.New("couldn't parse URL")
	}

	fullPath := parsedURL.Host + parsedURL.Path
	fullPath = strings.ToLower(fullPath)
	if fullPath == "" {
		return "", errors.New("couldn't parse URL")
	}

	if fullPath[len(fullPath)-1] == '/' {
		fullPath = fullPath[:len(fullPath)-1]
	}

	return fullPath, nil
}

/*
	if url[:5] == "https" && url[len(url)-1] == '/' {
		url := url[8 : len(url)-1]
		return url, nil
	} else if url[:4] == "http" && url[len(url)-1] == '/' {
		url := url[7 : len(url)-1]
		return url, nil
	} else if url[:5] == "https" {
		url := url[8:]
		return url, nil
	} else if url[:4] == "http" {
		url := url[7:]
		return url, nil
	}
*/
