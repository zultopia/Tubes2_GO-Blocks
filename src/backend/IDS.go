package main

import (
	"fmt"
)

func IDS(start, end WikiPage, maxDepth int) ([]WikiPage, int) {
	nodesChecked := 0
	cache := make(map[string][]WikiPage)
	for depth := 1; depth <= maxDepth; depth++ {
		path, nodesChecked := DLS(start, end, depth, &cache)
		// fmt.Println(path)
		if path != nil {
			return path, nodesChecked
		}
	}
	return nil, nodesChecked
}

// DLS up to a given depth
// var count = 0

func DLS(start, end WikiPage, depth int, cache *map[string][]WikiPage) ([]WikiPage, int) {
	if depth == 0 && start.Title != end.Title {
		return nil, 1
	}
	if start.Title == end.Title {
		return []WikiPage{start}, 1
	}
	currentChecked := 1
	var links []WikiPage
	if (*cache)[start.Title] == nil {
		links, _ = getWikiLinks(start, end)
		fmt.Printf("len before: ")
		fmt.Println(len(links))
		// count++
		(*cache)[start.Title] = links
	} else {
		links = (*cache)[start.Title]
		fmt.Println(len(links), depth)
	}
	// links, _ := getWikiLinks(start, end)
	fmt.Printf("len after: ")
	fmt.Println(len(links))
	fmt.Printf("depth: %d\n", depth)
	for i, link := range links {
		if link.Title == "Archimedes" {
			fmt.Println(i)
			fmt.Println(link.Title)
		}
		path, nodesChecked := DLS(link, end, depth-1, cache)
		currentChecked += nodesChecked
		// fmt.Println(path)
		if path != nil {
			return append([]WikiPage{start}, path...), currentChecked
		}
	}
	// fmt.Println(count)
	return nil, currentChecked
}
