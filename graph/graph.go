package graph

import (
	"fmt"
	"sync"
	"time"
)

type Graph interface {
	// Find returns all vertices connected to the given vertex
	Find(vertex string) []string
	// Similar to Find(), but performs non-blocking lookup
	Find2(vertex string) []string
	// Expands Find2, adds maximum thread count parameter
	Find3(vertex string, maxThread int) []string
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
	// lock queue and visited
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
		go func() {
			fmt.Println("looking up", el)
			children := g.lookup(el)

			mu.Lock()
			visited[el] = struct{}{}
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

// it is possible to look up the same vertex MULTIPLE times tho.
func (g graph) Find3(vertex string, tc int) []string {
	// lock queue and visited
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
		go func() {
			fmt.Println("looking up", el)
			children := g.lookup(el)

			mu.Lock()
			visited[el] = struct{}{}
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

func (g graph) lookup(vertex string) []string {
	time.Sleep(2 * time.Second)
	return g[vertex]
}
