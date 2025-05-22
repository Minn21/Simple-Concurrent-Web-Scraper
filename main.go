package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url>")
		os.Exit(1)
	}
	url := os.Args[1]

	links := make(chan string)
	var wg sync.WaitGroup

	// Fetch and parse in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Fetch error:", err)
			return
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Println("Parse error:", err)
			return
		}

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			link, exists := s.Attr("href")
			if exists {
				links <- link
			}
		})
	}()

	// Close the channel when done
	go func() {
		wg.Wait()
		close(links)
	}()

	// Create output file
	file, err := os.Create("links.txt")
	if err != nil {
		log.Fatal("Failed to create file:", err)
	}
	defer file.Close()

	// Write links to file
	for link := range links {
		_, err := file.WriteString(link + "\n")
		if err != nil {
			log.Println("Write error:", err)
		}
	}

	fmt.Println("Links saved to links.txt!")
}
