package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttr("a", "href")
		fmt.Println(links)
	})
	c.Visit("https://en.wikipedia.org/wiki/web_scraping")
}
