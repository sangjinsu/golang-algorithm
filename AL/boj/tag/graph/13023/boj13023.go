package main

import (
	"bufio"
	"fmt"
	"os"
)

var result = 0

func dfs(k, cnt int, visited []bool, graph map[int][]int) {
	if cnt == 4 {
		result = 1
		return
	}

	visited[k] = true
	for _, v := range graph[k] {
		if visited[v] == false {
			dfs(v, cnt+1, visited, graph)
		}
		if result == 1 {
			return
		}
	}
	visited[k] = false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	graph := make(map[int][]int)

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	visited := make([]bool, n)

	for i := 0; i < n && result != 1; i++ {
		dfs(i, 0, visited, graph)
	}

	fmt.Fprint(writer, result)

}
