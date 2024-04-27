package main

import (
	"fmt"
	"sync"
)

func IDS(start, end WikiPage, maxDepth int, multi bool) ([][]WikiPage, int) {
	nodesChecked := 0
	var solution [][]WikiPage
	var cache sync.Map
	for depth := 1; depth <= maxDepth; depth++ {
		/*
				path, nodesChecked := DLS(start, end, depth, &cache, &solution)
				// fmt.Println(path)
				if path != nil {
					return path, nodesChecked
				}
			}
			return nil, nodesChecked */
		var nodes int
		if (multi) {
			solution, nodes = DLSmulti(start, end, depth, &cache)
		} else {
			solution, nodes = DLSsingle(start, end, depth, &cache)
		}
		nodesChecked += nodes
		if len(solution) > 0 {
			break
		}
	}
	return solution, nodesChecked
}

// DLS up to a given depth
// var count = 0

func DLSmulti(start, end WikiPage, depth int, cache *sync.Map) ([][]WikiPage, int) {
	var max_go int = 15
	var guard = make(chan struct{}, max_go)
	var wg = sync.WaitGroup{}
	solution := make([][]WikiPage, 0)
	if depth == 0 && start.Title != end.Title {
		return nil, 1
	}
	if start.Title == end.Title {
		path := []WikiPage{start}
		solution = append(solution, path)
		return solution, 1
	}
	currentChecked := 1
	var links []WikiPage
	linkstmp, _ := cache.Load(start.Title)
	if linkstmp == nil {
		links, _ = getWikiLinks(start, end)
		cache.Store(start.Title, links)
	} else {
		links = linkstmp.([]WikiPage)
		fmt.Println(len(links), depth)
	}
	for _, link := range links {
		wg.Add(1)
		guard <- struct{}{}
		go func() {
			curSolution, nodesChecked := DLSmulti(link, end, depth-1, cache)
			currentChecked += nodesChecked
			if curSolution != nil {
				solution = append(solution, curSolution...)
			}
			<-guard
			defer wg.Done()
		}()
	}
	wg.Wait()

	for i, path := range solution {
		solution[i] = append([]WikiPage{start}, path...)
	}
	return solution, currentChecked
}

func DLSsingle(start, end WikiPage, depth int, cache *sync.Map) ([][]WikiPage, int) {
	var max_go int = 15
	var guard = make(chan struct{}, max_go)
	var wg = sync.WaitGroup{}
	solution := make([][]WikiPage, 0)
	if depth == 0 && start.Title != end.Title {
		return nil, 1
	}
	if start.Title == end.Title {
		path := []WikiPage{start}
		solution = append(solution, path)
		return solution, 1
	}
	currentChecked := 1
	var links []WikiPage
	linkstmp, _ := cache.Load(start.Title)
	if linkstmp == nil {
		links, _ = getWikiLinks(start, end)
		cache.Store(start.Title, links)
	} else {
		links = linkstmp.([]WikiPage)
	}
	for _, link := range links {
		wg.Add(1)
		guard <- struct{}{}
		go func() {
			curSolution, nodesChecked := DLSsingle(link, end, depth-1, cache)
			currentChecked += nodesChecked
			if curSolution != nil {
				solution = append(solution, curSolution...)
			}
			<-guard
			defer wg.Done()
		}()
		if (len(solution) != 0){
			solution[0] = append([]WikiPage{start}, solution[0]...)
			return solution, currentChecked
		}
	}
	wg.Wait()
	return solution, currentChecked
}