package main

import "fmt"

func dfs(graph [][]int, v int, visited []bool) {
	visited[v] = true
	fmt.Printf("%v ", v)
	for _, i := range graph[v] {
		if visited[i] != true {
			dfs(graph, i, visited)
		}
	}
}

func main() {
	graph := [][]int{
		{},
		{2, 3, 8},
		{1, 7},
		{1, 4, 5},
		{3, 5},
		{3, 4},
		{7},
		{2, 6, 8},
		{1, 7},
	}

	visited := make([]bool, 9)

	dfs(graph, 1, visited)
}
