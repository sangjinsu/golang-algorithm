package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var v, e int
	fmt.Fscanf(reader, "%d %d\n", &v, &e)

	parent := make([]int, v+1)
	for i := 1; i < v+1; i++ {
		parent[i] = i
	}

	var edges [][3]int
	for i := 0; i < e; i++ {
		var a, b, cost int
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &cost)
		edges = append(edges, [3]int{cost, a, b})
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	var mst [][2]int
	var result int
	for _, edge := range edges {
		cost, a, b := edge[0], edge[1], edge[2]
		if findParent(parent, a) != findParent(parent, b) {
			unionParent(parent, a, b)
			result += cost
			mst = append(mst, [2]int{a, b})
		}
	}

	fmt.Fprintln(writer, result, mst)
}
