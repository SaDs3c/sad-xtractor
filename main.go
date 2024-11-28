package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// Check if the user provided an argument
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./main <url>")
	}

	// Get the URL from the command-line argument
	url := os.Args[1]

	// Automatically add "http://" if no protocol is provided
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	// Create a new collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	// Slice to store the links
	var links []string

	// Callback for when an HTML element is visited
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		links = append(links, link)
	})

	// Callback for error handling
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error:", err)
	})

	// Start crawling
	fmt.Println("Visiting:", url)
	err := c.Visit(url)
	if err != nil {
		log.Fatal("Failed to visit the URL:", err)
	}

	// Print all the links found
	fmt.Println("\nLinks found:")
	for _, link := range links {
		fmt.Println("[+] "+ link + "\n")
	}
}
