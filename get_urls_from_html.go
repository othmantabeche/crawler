package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	var urls []string
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %w", err)
	}

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			return
		}

		href = strings.TrimSpace(href)
		relativeURL, err := url.Parse(href)

		if err != nil {
			fmt.Println("couldn't parse href: %w", err)
			return
		}

		resolved := baseURL.ResolveReference(relativeURL)
		urls = append(urls, resolved.String())
	})
	return urls, nil
}
