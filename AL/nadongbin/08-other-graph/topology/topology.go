package main

import (
	"container/list"
	"fmt"
)

func main() {
	var v, e int
	fmt.Scanln(&v, &e)

	indegree := make([]int, v+1)

	graph := make([][]int, v+1)

	for i := 0; i < e; i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		graph[a] = append(graph[a], b)
		indegree[b]++
	}

	topologySort(v, indegree, graph)

}

func topologySort(v int, indegree []int, graph [][]int) {
	result := []int{}
	queue := list.New()

	for i := 1; i < v+1; i++ {
		if indegree[i] == 0 {
			queue.PushBack(i)
		}
	}

	for queue.Len() > 0 {
		node := queue.Front()
		value := node.Value.(int)
		result = append(result, value)
		queue.Remove(node)

		for _, i := range graph[value] {
			indegree[i]--
			if indegree[i] == 0 {
				queue.PushBack(i)
			}
		}
	}

	for _, v := range result {
		fmt.Printf("%d ", v)
	}
}
