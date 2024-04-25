package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/gocolly/colly/v2"
)

// WikiPage represents a Wikipedia page with its title and URL
type WikiPage struct {
	Title string
	URL   string
}

var user_agent = []string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.5 Safari/605.1.15",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Windows; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/603.3.8 (KHTML, like Gecko) Version/10.1.2 Safari/603.3.8",
	"Mozilla/5.0 (Windows NT 10.0; Windows; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Safari/605.1.15",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Safari/605.1.15",
	"Mozilla/5.0 (Windows NT 10.0; Windows; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Safari/537.36"}

func getWikiLinks(page, end WikiPage) ([]WikiPage, error) {
	visited2 := make(map[string]bool)
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)
	// set fake Random User Agents
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", user_agent[rand.Intn(10)])
	})

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
