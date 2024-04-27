package main
import (
	"fmt"
	// "runtime/trace"
	"sync"
	"time"

	"github.com/orcaman/concurrent-map"
)
var wg = sync.WaitGroup{}
var max_go int = 100
var guard = make(chan struct{}, max_go)
var solution = make([][]WikiPage, 0)
// var level = 1

func BFSGo(start, end WikiPage) ([][]WikiPage, int) {
	if start.Title == end.Title{
		return [][]WikiPage{{end}}, 1
	}
	queue := make([][]WikiPage, 0)
	// var visited sync.Map
	visited2 := cmap.New()
	queue = append(queue, []WikiPage{start})
	newPath := make(chan []WikiPage)
	// visited := make(map[string]bool)
	// visited.Store(start.Title, true)
	visited2.Set(start.Title, true)
	go func() {
		defer close(newPath)
		for len(queue) > 0 {
			// fmt.Println(solF)
			tmpqueue := make([][]WikiPage, 0)
			// a := len(queue)
			for range queue {
				curpath := queue[0]
				queue = queue[1:]
				wg.Add(1)
				guard <- struct{}{}
				go BFSHelper(curpath, end, newPath, &visited2, &tmpqueue)
				// counter++
				// m.Lock()
				// tmpqueue = append(tmpqueue, <-newPath)
				// m.Unlock()
				// time.Sleep(5*time.Millisecond)
				// }
				// fmt.Println("Wait 5 sec")
				// time.Sleep(5 * time.Second)
			}
			wg.Wait()
			// level++
			// fmt.Println(len(queue))
			queue = tmpqueue
			// fmt.Println("Masuk Sini! level ", level)
			// fmt.Println(len(tmpqueue[0]))
			// fmt.Println(queue)
			// fmt.Println()
			time.Sleep(time.Second*2)
			if len(solution) > 0 {
				break
			}
		}
		// visited.Store(end, true)
		visited2.Set(end.Title, true)
		// fmt.Println("TESTES")
	}()
	for n := range newPath {
		path := n
		// fmt.Println(len(path))
		if len(path) == 0 {
			return nil, syncMapLen(&visited)
		}
		// if path[len(path)-1].Title == "Elon_Musk" {
		// 	fmt.Println(path)
		// }
		if path[len(path)-1].Title == end.Title {
			// fmt.Println(path)
			solution = append(solution, path)
			if !multi{
				// return solution, syncMapLen((&visited))
				return solution, visited2.Count()
			}
		}
	}
	// if solution tidak kosng, ....

	// return solution, syncMapLen(&visited)
	return solution, visited2.Count()
}
func BFSHelper(path []WikiPage, end WikiPage, newPath chan<- []WikiPage, visited2 *cmap.ConcurrentMap, tmpqueue *[][]WikiPage) {
	defer wg.Done()
	// time.Sleep(time.Second)
	if len(path) == 0 {
		newPath <- nil
		return
	}
	lastPage := path[len(path)-1]
	links, err := getWikiLinks(lastPage, end)
	// fmt.Println(len(links))
	if err != nil {
		// fmt.Println(path)
		newPath <- nil
		fmt.Println("error")
		return
	}
	// fmt.Println(1)
	for _, link := range links {
		// _, ok := visited.Load(link.Title)
		_, ok := visited2.Get(link.Title)
		if !ok {
			// fmt.Println(link.Title)
			if link.Title != end.Title {
				// visited.Store(link.Title, true)
				visited2.Set(link.Title, true)
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
