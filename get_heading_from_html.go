package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	for _, selector := range []string{"h1", "h2"} {
		heading := strings.TrimSpace(doc.Find(selector).First().Text())
		if heading != "" {
			return heading
		}
	}

	return ""
}
