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

	heading := doc.Find("h1, h2").Text()

	if len(heading) <= 0 {
		return ""
	}
	return heading
}
