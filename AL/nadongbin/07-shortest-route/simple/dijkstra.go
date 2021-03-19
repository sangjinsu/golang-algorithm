package main

import (
	"fmt"
	"math"
)

func getSmallestNode(n int, distance []int, visited []bool) int {
	minValue := math.MaxInt64
	index := 0
	for i := 1; i < n+1; i++ {
		if distance[i] < minValue && !visited[i] {
			minValue = distance[i]
			index = i
		}
	}
	return index
}
func dijkstra(n int, start int, graph [][][2]int, distance []int, visited []bool) {
	distance[start] = 0
	visited[start] = true
	for _, v := range graph[start] {
		distance[v[0]] = v[1]
	}

	for i := 0; i < n-1; i++ {
		now := getSmallestNode(n, distance, visited)
		visited[now] = true
		for _, v := range graph[now] {
			cost := distance[now] + v[1]
			if cost < distance[v[0]] {
				distance[v[0]] = cost
			}
		}
	}
}

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	var start int
	fmt.Scanln(&start)

	graph := make([][][2]int, n+1)
	visited := make([]bool, n+1)

	distance := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		distance[i] = math.MaxInt32
	}

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Scanln(&a, &b, &c)
		graph[a] = append(graph[a], [2]int{b, c})
	}

	fmt.Println(graph)

	dijkstra(n, start, graph, distance, visited)

	for i := 1; i < n+1; i++ {
		if distance[i] == math.MaxInt32 {
			fmt.Println("Infinity")
		} else {
			fmt.Println(distance[i])
		}
	}
}
