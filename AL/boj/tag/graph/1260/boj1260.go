package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func dfs(k int, graph map[int][]int, visited []bool) {
	visited[k] = true
	fmt.Fprintf(writer, "%d ", k)
	for _, v := range graph[k] {
		if visited[v] == false {
			dfs(v, graph, visited)
		}
	}
}

func bfs(k int, graph map[int][]int, visited []bool) {
	visited[k] = true
	queue := []int{k}
	for len(queue) > 0 {
		node := queue[0]
		fmt.Fprintf(writer, "%d ", node)
		queue = queue[1:]
		for _, v := range graph[node] {
			if visited[v] == false {
				queue = append(queue, v)
				visited[v] = true
			}
		}
	}
}

func main() {
	defer writer.Flush()

	var N, M, V int
	fmt.Fscanf(reader, "%d %d %d\n", &N, &M, &V)
	graph := make(map[int][]int)
	for i := 0; i < M; i++ {
		var v1, v2 int
		fmt.Fscanf(reader, "%d %d\n", &v1, &v2)
		graph[v1] = append(graph[v1], v2)
		graph[v2] = append(graph[v2], v1)
	}
	for _, ints := range graph {
		sort.Ints(ints)
	}

	visited := make([]bool, N+1)
	dfs(V, graph, visited)
	fmt.Fprintln(writer)

	visited = make([]bool, N+1)
	bfs(V, graph, visited)
}
