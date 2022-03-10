package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

var (
	V        int
	E        int
	K        int
	graph    [][][]int
	distance []int
)

type Item struct {
	dist int
	now  int
}

// An ItemHeap is a min-heap of Item.
type ItemHeap []Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Item))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d\n", &V, &E)
	fmt.Fscanf(reader, "%d\n", &K)

	graph = make([][][]int, V+1)

	for i := 0; i < E; i++ {
		var u, v, w int
		fmt.Fscanf(reader, "%d %d %d\n", &u, &v, &w)
		graph[u] = append(graph[u], []int{v, w})
	}

	distance = make([]int, V+1)
	for i := 0; i < V+1; i++ {
		distance[i] = math.MaxInt
	}

	dijkstra(K)
	for i := 1; i < V+1; i++ {
		dist := distance[i]
		if dist == math.MaxInt {
			fmt.Fprintln(writer, "INF")
		} else {
			fmt.Fprintln(writer, dist)
		}
	}
}

func dijkstra(start int) {
	queue := &ItemHeap{}
	heap.Init(queue)
	heap.Push(queue, Item{dist: 0, now: start})
	distance[start] = 0

	for queue.Len() > 0 {
		item := heap.Pop(queue).(Item)
		dist, now := item.dist, item.now
		if distance[now] < dist {
			continue
		}
		for _, adj := range graph[now] {
			cost := dist + adj[1]
			if cost < distance[adj[0]] {
				distance[adj[0]] = cost
				heap.Push(queue, Item{dist: cost, now: adj[0]})
			}
		}
	}
}
