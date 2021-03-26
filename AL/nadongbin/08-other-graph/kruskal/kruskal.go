package main

import (
	"fmt"
	"sort"
)

type edge struct {
	cost int
	a    int
	b    int
}

func main() {
	var v, e int
	fmt.Scanln(&v, &e)

	parent := make([]int, v+1)

	edges := make([]edge, e)
	result := 0

	for i := 1; i < len(parent); i++ {
		parent[i] = i
	}

	for i := 0; i < e; i++ {
		var a, b, cost int
		fmt.Scanln(&a, &b, &cost)
		edges[i] = edge{cost: cost, a: a, b: b}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	for _, edge := range edges {
		cost, a, b := edge.cost, edge.a, edge.b
		if findParent(parent, a) != findParent(parent, b) {
			unionParent(parent, a, b)
			result += cost
		}
	}

	fmt.Println(result)
	fmt.Println(parent)
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

func findParent(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = findParent(parent, parent[x])
	}
	return parent[x]
}
