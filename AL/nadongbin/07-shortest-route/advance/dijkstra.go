package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	distance int
	adjacent int
	index    int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue) Push(node interface{}) {
	n := len(*pq)
	item := node.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func dijkstra(n int, m int, start int, graph [][][2]int, distance []int) {
	pq := make(PriorityQueue, 1)
	pq[0] = &Node{distance: 0, adjacent: start, index: 0}
	heap.Init(&pq)
	distance[start] = 0
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)

		dist, now := node.distance, node.adjacent
		if distance[now] < dist {
			continue
		}
		for _, v := range graph[now] {
			cost := dist + v[1]
			if cost < distance[v[0]] {
				distance[v[0]] = cost
				heap.Push(&pq, &Node{distance: cost, adjacent: v[0]})
			}
		}
	}
}

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	var start int
	fmt.Scanln(&start)

	graph := make([][][2]int, n+1)

	distance := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		distance[i] = math.MaxInt32
	}

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Scanln(&a, &b, &c)
		graph[a] = append(graph[a], [2]int{b, c})
	}

	fmt.Println(graph)

	dijkstra(n, m, start, graph, distance)

	for i := 1; i < n+1; i++ {
		if distance[i] == math.MaxInt32 {
			fmt.Println("Infinity")
		} else {
			fmt.Println(distance[i])
		}
	}
}
