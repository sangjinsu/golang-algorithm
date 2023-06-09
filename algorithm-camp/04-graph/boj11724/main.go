package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(graph [][]int, visited []bool, v int) {
	visited[v] = true
	for _, vertex := range graph[v] {
		if !visited[vertex] {
			dfs(graph, visited, vertex)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	graph := make([][]int, N+1)
	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscanf(reader, "%d %d\n", &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var result int
	visited := make([]bool, N+1)
	for i := 1; i < N+1; i++ {
		if !visited[i] {
			result++
			dfs(graph, visited, i)
		}
	}

	fmt.Fprint(writer, result)
}
