package main

import (
	"fmt"
	"sort"
)

func findParent(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = findParent(parent, parent[x])
	}
	return parent[x]
}

func unionParent(parent []int, a, b int) {
	a = findParent(parent, a)
	b = findParent(parent, b)
	if a < b {
		parent[b] = a
	} else {
		parent[a] = b
	}
}

type edge struct {
	from int
	to   int
	cost int
}

func main() {
	var v, e int
	fmt.Scanln(&v, &e)
	parent := make([]int, v+1)

	edges := []edge{}
	result := 0

	for i := 1; i < v+1; i++ {
		parent[i] = i
	}

	for i := 0; i < e; i++ {
		var a, b, cost int
		fmt.Scanln(&a, &b, &cost)
		edges = append(edges, edge{from: a, to: b, cost: cost})
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	last := 0

	for _, edge := range edges {
		cost, a, b := edge.cost, edge.from, edge.to
		if findParent(parent, a) != findParent(parent, b) {
			unionParent(parent, a, b)
			result += cost
			last = cost
		}
	}

	fmt.Println(result - last)
}
