package graph

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Graph interface {
	// Find returns all vertices connected to the given vertex
	Find(vertex string) []string
	// Similar to Find(), but performs non-blocking lookup
	Find2(vertex string) []string
	// Expands Find2, adds maximum thread count parameter
	Find3(vertex string, maxThread int) []string
	// Same as Find3 but use a Cond variable instead of WaitGroup/buffered channel
	Find4(vertex string, maxThread int) []string
}

type graph map[string][]string

func DefaultGraph() Graph {
	g := make(graph)
	g["a"] = []string{"b", "c", "d"}
	g["b"] = []string{"c"}
	g["c"] = []string{"c", "d"}
	g["d"] = []string{"a", "c"}
	g["e"] = []string{} // e is forever alone
	return g
}

// traverse the graph using BFS
func (g graph) Find(vertex string) []string {
	var queue []string
	visited := make(map[string]struct{})

	queue = append(queue, vertex)

	for len(queue) != 0 {
		// pop
		el := queue[0]
		queue = queue[1:]

		if _, ok := visited[el]; ok {
			continue
		}

		visited[el] = struct{}{}
		children := g.lookup(el)
		queue = append(queue, children...)
	}

	var result []string
	for v, _ := range visited {
		result = append(result, v)
	}
	return result
}

func (g graph) Find2(vertex string) []string {
	// lock queue
	var mu sync.Mutex
	var wg sync.WaitGroup

	var queue []string
	visited := make(map[string]struct{})

	queue = append(queue, vertex)

	for {
		if len(queue) == 0 {
			wg.Wait()
		}

		if len(queue) == 0 {
			break
		}

		// pop
		mu.Lock()
		el := queue[0]
		queue = queue[1:]
		mu.Unlock()

		if _, ok := visited[el]; ok {
			continue
		}

		wg.Add(1)
		visited[el] = struct{}{}
		go func() {
			fmt.Println("looking up", el)
			children := g.lookup(el)

			mu.Lock()
			queue = append(queue, children...)
			mu.Unlock()
			wg.Done()
		}()
	}

	var result []string
	for v, _ := range visited {
		result = append(result, v)
	}
	return result
}

func (g graph) Find3(vertex string, tc int) []string {
	// lock queue
	var mu sync.Mutex
	var wg sync.WaitGroup
	sem := make(chan struct{}, tc)

	var queue []string
	visited := make(map[string]struct{})

	queue = append(queue, vertex)

	for {
		if len(queue) == 0 {
			wg.Wait()
		}

		if len(queue) == 0 {
			break
		}

		fmt.Println(queue)
		mu.Lock()
		el := queue[0]
		queue = queue[1:]
		mu.Unlock()

		if _, ok := visited[el]; ok {
			continue
		}

		sem <- struct{}{}

		wg.Add(1)
		visited[el] = struct{}{}
		go func() {
			fmt.Println("looking up", el)
			children := g.lookup(el)

			mu.Lock()
			queue = append(queue, children...)
			mu.Unlock()
			wg.Done()
			<-sem
		}()
	}

	var result []string
	for v, _ := range visited {
		result = append(result, v)
	}
	return result
}

func (g graph) Find4(vertex string, tc int) []string {
	var mu sync.Mutex
	var queue []string
	queue = append(queue, vertex)

	var threads int32
	cond := sync.Cond{L: &sync.Mutex{}}

	visited := make(map[string]struct{})

	for {
		if len(queue) == 0 {
			cond.L.Lock()
			for atomic.LoadInt32(&threads) > 0 {
				cond.Wait()
			}
			cond.L.Unlock()
		}

		if len(queue) == 0 {
			break
		}

		mu.Lock()
		el := queue[0]
		queue = queue[1:]
		mu.Unlock()

		if _, ok := visited[el]; ok {
			continue
		}

		// check for thread count
		cond.L.Lock()
		for atomic.LoadInt32(&threads) == int32(tc) {
			cond.Wait()
		}
		cond.L.Unlock()

		visited[el] = struct{}{}
		atomic.AddInt32(&threads, 1)
		go func() {
			fmt.Println("looking up", el)
			children := g.lookup(el)

			mu.Lock()
			queue = append(queue, children...)
			mu.Unlock()

			atomic.AddInt32(&threads, -1)
			cond.Signal()
		}()
	}

	var result []string
	for v, _ := range visited {
		result = append(result, v)
	}
	return result
}

func (g graph) lookup(vertex string) []string {
	time.Sleep(2 * time.Second)
	return g[vertex]
}
