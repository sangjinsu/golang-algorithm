package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	adjacent int
	distance int
	index    int
}

type PriorityQueue []*Edge

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

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Edge)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*pq = old[:n-1]
	return node
}

func dijkstra(start int, graph [][]Edge, distance []int) {
	pq := make(PriorityQueue, 1)
	pq[0] = &Edge{adjacent: start, distance: 0, index: 0}
	distance[start] = 0

	heap.Init(&pq)

	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Edge)
		now, d := v.adjacent, v.distance

		if distance[now] < d {
			continue
		}

		for _, v := range graph[now] {
			cost := d + v.distance
			if cost < distance[v.adjacent] {
				distance[v.adjacent] = cost
				heap.Push(&pq, &Edge{adjacent: v.adjacent, distance: cost})
			}
		}
	}
}

func main() {
	var N, M, C int
	fmt.Scanln(&N, &M, &C)

	graph := make([][]Edge, N+1)

	for i := 0; i < M; i++ {
		var X, Y, Z int
		fmt.Scanln(&X, &Y, &Z)
		graph[X] = append(graph[X], Edge{adjacent: Y, distance: Z})
	}

	distance := make([]int, N+1)

	for i := 0; i < N+1; i++ {
		distance[i] = math.MaxInt32
	}

	dijkstra(C, graph, distance)

	visitedCity := 0
	maxTime := 0
	for _, d := range distance {
		if d != math.MaxInt32 {
			visitedCity++
			maxTime = int(math.Max(float64(maxTime), float64(d)))
		}
	}

	fmt.Println(distance)
	fmt.Println(visitedCity-1, maxTime)

}
