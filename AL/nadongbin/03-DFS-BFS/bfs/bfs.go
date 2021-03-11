package main

import (
	"container/list"
	"fmt"
)

func bfs(graph [][]int, v int, visited []bool) {
	queue := list.New()
	queue.PushBack(v)
	visited[v] = true

	for queue.Len() > 0 {
		v := queue.Front()
		queue.Remove(v)
		fmt.Printf("%v ", v.Value)

		for _, i := range graph[v.Value.(int)] {
			if visited[i] == false {
				queue.PushBack(i)
				visited[i] = true
			}
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

	bfs(graph, 1, visited)
}
