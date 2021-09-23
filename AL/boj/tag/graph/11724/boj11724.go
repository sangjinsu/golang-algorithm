package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(k int, graph map[int][]int, visited []bool) {
	visited[k] = true
	for _, v := range graph[k] {
		if visited[v] == false {
			dfs(v, graph, visited)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)
	graph := make(map[int][]int, N+1)

	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscanf(reader, "%d %d\n", &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var cnt int
	visited := make([]bool, N+1)
	for i := 1; i < N+1; i++ {
		if visited[i] == false {
			dfs(i, graph, visited)
			cnt++
		}
	}

	fmt.Fprint(writer, cnt)
}
