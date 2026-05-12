package main

import (
	"fmt"
	"os"
)

func main() {
	comandLineArgs := os.Args[1:]

	if len(comandLineArgs) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(comandLineArgs) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := comandLineArgs[0]

	if len(rawBaseURL) == 1 {
		fmt.Println("starting crawl of:", rawBaseURL)
	}

	pages := make(map[string]int)

	crawlPage(rawBaseURL, rawBaseURL, pages)

	for page, count := range pages {
		fmt.Printf("%v -- %v\n", page, count)
	}

	/*
		html, err := getHTML(comandLineArgs[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(html)
	*/
}
