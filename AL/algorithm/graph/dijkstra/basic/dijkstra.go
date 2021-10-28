package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdin)
	defer writer.Flush()

	var N, E int
	fmt.Fscanf(reader, "%d %d\n", &N, &E)
	G := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		G[i] = make([]int, N+1)
	}
	D := make([][2]int, N+1)
	for i := 0; i < N+1; i++ {
		D[i] = [2]int{math.MaxInt64, 0}
	}
	for i := 0; i < E; i++ {
		var s1, s2, w int
		fmt.Fscanf(reader, "%d %d %d\n", &s1, &s2, &w)
		G[s1][s2] = w
	}

	dijkstra(N, G, D)
	fmt.Fprintln(writer, D)
}

func dijkstra(N int, G [][]int, D [][2]int) {
	curV := 0
	visited := make([]bool, N+1)
	U := []int{curV}
	visited[curV] = true
	for i := 0; i < N+1; i++ {
		if G[curV][i] > 0 {
			D[i][0] = G[curV][i]
		}
	}

	for len(U) <= N {
		minV := math.MaxInt64
		for i := 0; i < N+1; i++ {
			if visited[i] {
				continue
			}
			if D[i][0] < minV {
				minV = D[i][0]
				curV = i
			}
		}

		U = append(U, curV)

	}
}
