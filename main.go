package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	c.OnHTML(".mw-parser-output", func(h *colly.HTMLElement) {
		links := h.ChildAttrs("a", "href")
		fmt.Println(links)
		// fmt.Println(h.Text)
	})
	c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
}
