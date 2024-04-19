package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// WikiPage represents a Wikipedia page with its title and URL
type WikiPage struct {
	Title string
	URL   string
}

// getWikiLinks without cache
func getWikiLinks(page WikiPage) []WikiPage {
	links := make([]WikiPage, 0)
	resp, err := http.Get(page.URL)
	if err != nil {
		fmt.Println("Error fetching page:", err)
		return links
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" && strings.HasPrefix(attr.Val, "/wiki/") {
						title := strings.TrimPrefix(attr.Val, "/wiki/")
						link := WikiPage{Title: title, URL: "https://en.wikipedia.org" + attr.Val}
						links = append(links, link)
					}
				}
			}
		}
	}
}

// BFS Algorithm
func BFS(start, end WikiPage) ([]WikiPage, int) {
	queue := make([][]WikiPage, 0)
	queue = append(queue, []WikiPage{start})
	visited := make(map[string]bool)
	visited[start.Title] = true

	for len(queue) > 0 {
		// Dequeue
		path := queue[0]
		queue = queue[1:]

		// Ambil page yang terakhir dicek
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
	nodesChecked := 0
	for depth := 1; depth <= maxDepth; depth++ {
		path, nodesChecked := DLS(start, end, depth)
		if path != nil {
			return path, nodesChecked
		}
	}
	return nil, nodesChecked
}

// DLS up to a given depth
func DLS(start, end WikiPage, depth int) ([]WikiPage, int) {
	if depth == 0 && start.Title != end.Title {
		return nil, 1
	}
	if start.Title == end.Title {
		return []WikiPage{start}, 1
	}
	currentChecked := 1
	links := getWikiLinks(start)
	for _, link := range links {
		path, nodesChecked := DLS(link, end, depth-1)
		currentChecked += nodesChecked
		if path != nil {
			return append([]WikiPage{start}, path...), currentChecked
		}
	}
	return nil, currentChecked
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
