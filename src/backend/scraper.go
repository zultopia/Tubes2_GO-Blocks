package main

import (
	"fmt"
	"strings"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

// WikiPage represents a Wikipedia page with its title and URL
type WikiPage struct {
	Title string
	URL   string
}

func getWikiLinks(page, end WikiPage) ([]WikiPage, error) {
	visited2 := make(map[string]bool)
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)
	// Add Random User Agents
	extensions.RandomUserAgent(c)

	var wikipages []WikiPage
	var wikipage WikiPage
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		tmp := e.Attr("href")
		// fmt.Println(tmp)
		if strings.HasPrefix(tmp, "/wiki") && !strings.HasPrefix(tmp, "/wiki/File:") && !strings.HasPrefix(tmp, "#") && !strings.HasPrefix(tmp, "https://") && !strings.HasPrefix(tmp, "/wiki/Special:") && !strings.HasPrefix(tmp, "/wiki/Category:") {
			wikipage.URL = "https://en.wikipedia.org" + tmp
			// fmt.Println(wikipage.URL)
			wikipage.Title = strings.TrimPrefix(wikipage.URL, "https://en.wikipedia.org/wiki/")
			// fmt.Println(wikipage.Title)
			if !visited2[wikipage.Title] {
				wikipages = append(wikipages, wikipage)
				visited2[wikipage.Title] = true
			}
		}
		// time.Sleep(5 * time.Millisecond)
		if wikipage.Title == end.Title {
			return
		}

	})
	err := c.Visit(page.URL)
	if err != nil {
		return nil, err
	}
	return wikipages, err
}
