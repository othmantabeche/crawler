package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	var urls []string
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %w", err)
	}

	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok {
			return
		}

		src = strings.TrimSpace(src)
		relativeURL, err := url.Parse(src)
		if err != nil {
			fmt.Println("couldn't parse HTML: %w", err)
			return
		}

		resolved := baseURL.ResolveReference(relativeURL)
		urls = append(urls, resolved.String())
	})

	return urls, nil
}
