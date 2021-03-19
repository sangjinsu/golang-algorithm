package main

import (
	"fmt"
	"math"
)

const Inf int = math.MaxInt32

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	graph := make([][]int, n+1)

	row := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		row[i] = Inf
	}

	for i := 0; i < n+1; i++ {
		r := make([]int, n+1)
		copy(r, row)
		graph[i] = r
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		graph[a][b] = 1
		graph[b][a] = 1
	}

	var x, k int
	fmt.Scanln(&x, &k)

	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == j {
				graph[i][j] = 0
			}
		}
	}

	for k := 1; k < n+1; k++ {
		for i := 1; i < n+1; i++ {
			for j := 1; j < n+1; j++ {
				graph[i][j] = int(math.Min(float64(graph[i][j]), float64(graph[i][k]+graph[k][j])))
			}
		}
	}

	result := graph[1][k] + graph[k][x]

	if result >= Inf {
		fmt.Println(-1)
	} else {
		fmt.Println(result)
	}

}
