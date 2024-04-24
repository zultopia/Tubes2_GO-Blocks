package main

import (
	"fmt"
	"sync"
)

func IDS(start, end WikiPage, maxDepth int) ([][]WikiPage, int) {
	nodesChecked := 0
	var solution [][]WikiPage
	cache := make(map[string][]WikiPage)
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
		solution, nodes = DLS(start, end, depth, &cache)
		nodesChecked += nodes
		if len(solution) > 0 {
			break
		}
	}
	return solution, nodesChecked
}

// DLS up to a given depth
// var count = 0

func DLS(start, end WikiPage, depth int, cache *map[string][]WikiPage) ([][]WikiPage, int) {
	var max_go int = 100
	var guard = make(chan struct{}, max_go)
	var wg = sync.WaitGroup{}
	solution := make([][]WikiPage, 0)
	if depth == 0 && start.Title != end.Title {
		return nil, 1
	}
	if start.Title == end.Title {
		/*
			return []WikiPage{start}, 1
		*/
		path := []WikiPage{start}
		solution = append(solution, path)
		return solution, 1
	}
	currentChecked := 1
	var links []WikiPage
	if (*cache)[start.Title] == nil {
		links, _ = getWikiLinks(start, end)
		// fmt.Printf("len before: ")
		// fmt.Println(len(links))
		// count++
		(*cache)[start.Title] = links
	} else {
		links = (*cache)[start.Title]
		fmt.Println(len(links), depth)
	}
	// links, _ := getWikiLinks(start, end)
	// fmt.Printf("len after: ")
	// fmt.Println(len(links))
	// fmt.Printf("depth: %d\n", depth)
	for _, link := range links {
		// if link.Title == "Archimedes" {
		// 	fmt.Println(i)
		// 	fmt.Println(link.Title)
		// }
		/*
				path, nodesChecked := DLS(link, end, depth-1, cache, solution)
				currentChecked += nodesChecked
				// fmt.Println(path)
				if path != nil {
					return append([]WikiPage{start}, path...), currentChecked
				}
			}
			// fmt.Println(count)
			return nil, currentChecked */
		// fmt.Println(link)
		wg.Add(1)
		guard <- struct{}{}
		go func(){
			curSolution, nodesChecked := DLS(link, end, depth-1, cache)
			currentChecked += nodesChecked
			if curSolution != nil {
				solution = append(solution, curSolution...)
			}
			<- guard
			defer wg.Done()
		}()
	}
	wg.Wait()


	for i, path := range solution {
		solution[i] = append([]WikiPage{start}, path...)
	}
	return solution, currentChecked
}
