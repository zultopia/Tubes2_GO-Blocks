package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var max_go int = 150
var guard = make(chan struct{}, max_go)
var solution = make([][]WikiPage, 0)
var m = sync.RWMutex{}

func BFSGo(start, end WikiPage, multi bool) ([][]WikiPage, int) {
	solution = make([][]WikiPage, 0)
	wg = sync.WaitGroup{}
	guard = make(chan struct{}, max_go)
	m = sync.RWMutex{}
	if start.Title == end.Title {
		return [][]WikiPage{{end}}, 1
	}
	queue := make([][]WikiPage, 0)
	var visited sync.Map
	queue = append(queue, []WikiPage{start})
	newPath := make(chan []WikiPage)
	visited.Store(start.Title, true)
	go func() {
		defer close(newPath)
		for len(queue) > 0 {
			tmpqueue := make([][]WikiPage, 0)
			for _, curpath := range queue {
				wg.Add(1)
				guard <- struct{}{}
				go BFSHelper(curpath, end, newPath, &visited, &tmpqueue)
			}
			wg.Wait()
			queue = tmpqueue
			time.Sleep(time.Second * 2)
			if len(solution) > 0 {
				break
			}
		}
		visited.Store(end, true)
	}()
	for n := range newPath {
		path := n
		if path == nil {
			continue
		}
		if path[len(path)-1].Title == end.Title {
			fmt.Println(path)
			solution = append(solution, path)
			if !multi {
				return solution, syncMapLen((&visited))
			}
		}
	}

	return solution, syncMapLen(&visited)
}
func BFSHelper(path []WikiPage, end WikiPage, newPath chan<- []WikiPage, visited *sync.Map, tmpqueue *[][]WikiPage) {
	defer wg.Done()
	if len(path) == 0 {
		newPath <- nil
		return
	}
	lastPage := path[len(path)-1]
	links, err := getWikiLinks(lastPage, end)
	if err != nil {
		newPath <- nil
		fmt.Println("error")
		return
	}
	for _, link := range links {
		_, ok := visited.Load(link.Title)
		if !ok {
			if link.Title != end.Title {
				visited.Store(link.Title, true)
			}
			newPathtmp := append([]WikiPage{}, path...)
			newPathtmp = append(newPathtmp, link)
			newPath <- newPathtmp
			m.Lock()
			*tmpqueue = append(*tmpqueue, newPathtmp)
			m.Unlock()
		}
	}
	<-guard
}

func syncMapLen(sm *sync.Map) int {
	var i int
	sm.Range(func(k, v interface{}) bool {
		i++
		return true
	})
	return i
}
