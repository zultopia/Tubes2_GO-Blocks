package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// WikiPage represents a Wikipedia page with its title and URL
type WikiPage struct {
	Title string
	URL   string
}

var WikiPageCache sync.Map

// BFS Algorithm
func BFS(start, end WikiPage) ([]WikiPage, int) {
	queue := make([][]WikiPage, 0)
	queue = append(queue, []WikiPage{start})
	visited := make(map[string]bool)
	visited[start.Title] = true

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		lastPage := path[len(path)-1]
		if lastPage.Title == end.Title {
			return path, len(visited)
		}

		links := getWikiLinks(lastPage)
		for _, link := range links {
			if !visited[link.Title] {
				visited[link.Title] = true
				newPath := append([]WikiPage{}, path...)
				newPath = append(newPath, link)
				queue = append(queue, newPath)
			}
		}
	}

	return nil, len(visited)
}

// IDS Algorithm
func IDS(start, end WikiPage, maxDepth int) ([]WikiPage, int) {
	for depth := 1; depth <= maxDepth; depth++ {
		path, nodesChecked := DLS(start, end, depth)
		if path != nil {
			return path, nodesChecked
		}
	}
	return nil, -1
}

// DLS up to a given depth 
func DLS(start, end WikiPage, depth int) ([]WikiPage, int) {
	if depth == 0 && start.Title != end.Title {
		return nil, 1
	}
	if start.Title == end.Title {
		return []WikiPage{start}, 1
	}

	links := getWikiLinks(start)
	for _, link := range links {
		path, nodesChecked := DLS(link, end, depth-1)
		if path != nil {
			return append([]WikiPage{start}, path...), nodesChecked
		}
	}
	return nil, len(links)
}

// getWikiLinks with cache
func getWikiLinks(page WikiPage) []WikiPage {
	// Cek cache, jika sudah ada tautan sebelumnya, kembalikan dari cache
	if cachedLinks, ok := WikiPageCache.Load(page.URL); ok {
		if links, ok := cachedLinks.([]WikiPage); ok {
			// fmt.Println("Fetched links from cache:", page.Title)
			return links
		}
	}

	links := make([]WikiPage, 0)
	doc, err := goquery.NewDocument(page.URL)
	if err != nil {
		fmt.Println("Error fetching page:", err)
		return links
	}

	doc.Find("#mw-content-text a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.HasPrefix(href, "/wiki/") {
			title := strings.TrimPrefix(href, "/wiki/")
			link := WikiPage{Title: title, URL: "https://en.wikipedia.org" + href}
			links = append(links, link)
		}
	})
	// Simpan hasil parsing ke cache
	WikiPageCache.Store(page.URL, links)

	return links
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run wikirace.go <algorithm> <start_article> <end_article>")
		return
	}

	algorithm := os.Args[1]
	start := WikiPage{Title: os.Args[2], URL: "https://en.wikipedia.org/wiki/" + os.Args[2]}
	end := WikiPage{Title: os.Args[3], URL: "https://en.wikipedia.org/wiki/" + os.Args[3]}

	var path []WikiPage
	var nodesChecked int
	startTime := time.Now()

	switch algorithm {
	case "BFS":
		path, nodesChecked = BFS(start, end)
	case "IDS":
		path, nodesChecked = IDS(start, end, 20) // Maximum depth for IDS
	default:
		fmt.Println("Invalid algorithm. Please use 'BFS' or 'IDS'.")
		return
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	if path != nil {
		fmt.Printf("Number of articles checked: %d\n", nodesChecked)
		fmt.Printf("Number of articles traversed: %d\n", len(path))
		fmt.Println("Traversal route:")
		for _, page := range path {
			fmt.Printf("- %s\n", page.Title)
		}
		fmt.Printf("Search time: %d ms\n", elapsedTime.Milliseconds())
	} else {
		fmt.Println("No path found.")
	}
}