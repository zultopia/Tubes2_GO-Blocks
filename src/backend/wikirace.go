package main

import (
	"fmt"
	// "net/http"
	"os"
	"strings"
	"sync"
	"time"
	"github.com/gocolly/colly/v2"
	// "golang.org/x/net/html"
)

// WikiPage represents a Wikipedia page with its title and URL
type WikiPage struct {
	Title string
	URL   string
}

// getWikiLinks without cache
// func getWikiLinks(page WikiPage) []WikiPage {
// 	links := make([]WikiPage, 0)
// 	resp, err := http.Get(page.URL)
// 	if err != nil {
// 		fmt.Println("Error fetching page:", err)
// 		return links
// 	}
// 	defer resp.Body.Close()

// 	z := html.NewTokenizer(resp.Body)
// 	for {
// 		tt := z.Next()
// 		switch tt {
// 		case html.ErrorToken:
// 			return links
// 		case html.StartTagToken, html.SelfClosingTagToken:
// 			t := z.Token()
// 			if t.Data == "a" {
// 				for _, attr := range t.Attr {
// 					if attr.Key == "href" && strings.HasPrefix(attr.Val, "/wiki/") {
// 						title := strings.TrimPrefix(attr.Val, "/wiki/")
// 						link := WikiPage{Title: title, URL: "https://en.wikipedia.org" + attr.Val}
// 						links = append(links, link)
// 					}
// 				}
// 			}
// 		}
// 	}
// }

func getWikiLinks(page WikiPage) ([]WikiPage, error){
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)
	var wikipages []WikiPage
	var wikipage WikiPage
	c.OnHTML("a[href]", func(e *colly.HTMLElement){
		tmp := e.Attr("href")
		// fmt.Println(tmp)
		if strings.HasPrefix(tmp, "/wiki") && !strings.HasPrefix(tmp, "/wiki/File:"){
			wikipage.URL = "https://en.wikipedia.org" + tmp
			// fmt.Println(wikipage.URL)
			wikipage.Title = strings.TrimPrefix(wikipage.URL, "https://en.wikipedia.org/wiki/")
			// fmt.Println(wikipage.Title)
			wikipages = append(wikipages, wikipage)
			// time.Sleep(5 * time.Millisecond)
		}
		
	})
	err := c.Visit(page.URL)
	if err != nil{
		return nil, err
	}
	return wikipages, err
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

		links, _ := getWikiLinks(lastPage)
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
var m = sync.RWMutex{}
// var m2 = sync.RWMutex{}
var wg = sync.WaitGroup{}
func BFSGo(start, end WikiPage) ([]WikiPage, int) {
	queue := make([][]WikiPage, 0)
	queue = append(queue, []WikiPage{start})
	newPath := make(chan []WikiPage)
	var visited sync.Map
	// visited := make(map[string]bool)
	visited.Store(start.Title, true)
	go func(){
		defer close(newPath)
		for len(queue) > 0{
			curpath := queue[0]
			queue = queue[1:]
			wg.Add(1)
			go BFSHelper(curpath, newPath, &visited)
			m.Lock()
			queue = append(queue, <-newPath)
			m.Unlock()
		}
		wg.Wait()
	}()
	for n := range newPath {
		path := n
		if len(path) == 0{
			return nil, syncMapLen(&visited)
		}
 		// fmt.Println(path)
		if path[len(path) - 1].Title == end.Title {
			return path, syncMapLen(&visited)
		}
		
		
	}
	return nil, syncMapLen(&visited)
}
func BFSHelper(path []WikiPage, newPath chan <-[]WikiPage, visited *sync.Map){
	defer wg.Done()
	if len(path) == 0{
		newPath <- []WikiPage{}
		return
	}
	lastPage := path[len(path)-1]
	links, err := getWikiLinks(lastPage)
	fmt.Println(len(links))
	if err != nil{
		newPath <- []WikiPage{}
		return
	}
	// fmt.Println(1)
	for _, link := range links {
		_, ok := visited.Load(link.Title)
		if !ok {
			visited.Store(link.Title, true)
			newPathtmp := append([]WikiPage{}, path...)
			newPathtmp = append(newPathtmp, link)
			newPath <- newPathtmp
			
		}
	}

}
func syncMapLen(sm *sync.Map) int {
	var i int
	sm.Range(func(k, v interface{}) bool {
        i++
        return true
    })
    return i
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
	links, _ := getWikiLinks(start)
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
		path, nodesChecked = BFSGo(start, end)
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
