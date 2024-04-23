package main

import (
)

func IDS(start, end WikiPage, maxDepth int) ([]WikiPage, int) {
	nodesChecked := 0
	var cache map[WikiPage][]WikiPage
	for depth := 1; depth <= maxDepth; depth++ {
		path, nodesChecked := DLS(start, end, depth, cache)
		if path != nil {
			return path, nodesChecked
		}
	}
	return nil, nodesChecked
}

// DLS up to a given depth
func DLS(start, end WikiPage, depth int, cache map[WikiPage][]WikiPage) ([]WikiPage, int) {
	if depth == 0 && start.Title != end.Title {
		return nil, 1
	}
	if start.Title == end.Title {
		return []WikiPage{start}, 1
	}
	currentChecked := 1
	links := []WikiPage{}
	if cache[start] == nil{
		links, _ := getWikiLinks(start, end)
		cache[start] = links 
	}else{
		links = cache[start] 
	}
	for _, link := range links {
		path, nodesChecked := DLS(link, end, depth-1, cache)
		currentChecked += nodesChecked
		if path != nil {
			return append([]WikiPage{start}, path...), currentChecked
		}
	}
	return nil, currentChecked
}