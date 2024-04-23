package main

import (
	"fmt"
)

func IDS(start, end WikiPage, maxDepth int) ([]WikiPage, int) {
	maxDepth = 3
	nodesChecked := 0
	cache := make(map[string][]WikiPage)
	for depth := 1; depth <= maxDepth; depth++ {
		path, nodesChecked := DLS(start, end, depth, &cache)
		fmt.Println(depth)
		if path != nil {
			return path, nodesChecked
		}
	}
	return nil, nodesChecked
}

// DLS up to a given depth
// var count = 0
func DLS(start, end WikiPage, depth int, cache *map[string][]WikiPage) ([]WikiPage, int) {
	if start.Title == end.Title {
		return []WikiPage{start}, 1
	}
	if depth == 0{
		return nil, 1
	}
	currentChecked := 1
	links := []WikiPage{}
	if (*cache)[start.Title] == nil{
		links, _ := getWikiLinks(start, end)
		// count++
		(*cache)[start.Title] = links 
	}else{
		links = (*cache)[start.Title] 
	}
	for _, link := range links {
		path, nodesChecked := DLS(link, end, depth-1, cache)
		currentChecked += nodesChecked
		if path != nil {
			return append([]WikiPage{start}, path...), currentChecked
		}
	}
	// fmt.Println(count)
	return nil, currentChecked
}