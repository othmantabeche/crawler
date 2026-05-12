package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parseRawBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("Fail parsing URL")
		return
	}

	parseRawCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Fail parsing URL")
		return
	}

	if parseRawBaseURL.Hostname() != parseRawCurrentURL.Hostname() {
		return
	}

	normalizeCurrentUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Fail normalizing URL")
		return
	}

	if _, visited := pages[normalizeCurrentUrl]; visited {
		pages[normalizeCurrentUrl]++
		return
	}

	pages[normalizeCurrentUrl] = 1
	fmt.Printf("crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("Fail geting the HTML")
		return
	}

	getUrls, err := getURLsFromHTML(html, parseRawCurrentURL)
	if err != nil {
		fmt.Println("Fail geting the URL`s")
		return
	}

	for _, nextURL := range getUrls {
		crawlPage(rawBaseURL, nextURL, pages)
	}

}
