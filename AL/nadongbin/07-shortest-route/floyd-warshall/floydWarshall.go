package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	graph := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			graph[i] = append(graph[i], math.MaxInt32)
		}
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == j {
				graph[i][j] = 0
			}
		}
	}

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Scanln(&a, &b, &c)
		graph[a][b] = c
	}

	for k := 1; k < n+1; k++ {
		for i := 1; i < n+1; i++ {
			for j := 1; j < n+1; j++ {
				graph[i][j] = int(math.Min(float64(graph[i][j]), float64(graph[i][k]+graph[k][j])))
			}
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if graph[i][j] == math.MaxInt32 {
				fmt.Printf("Infinity ")
			} else {
				fmt.Printf("%d ", graph[i][j])
			}
		}
		fmt.Println()
	}
}
