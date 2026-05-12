package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	selection := doc.Find("main p").First()

	if selection.Length() == 0 {
		selection = doc.Find("p").First()
	}

	paragraph := strings.TrimSpace(selection.Text())

	if len(paragraph) <= 0 {
		return ""
	}
	return paragraph
}
