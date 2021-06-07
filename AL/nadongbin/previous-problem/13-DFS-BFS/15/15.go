package main

import (
	"container/list"
	"fmt"
)

func main() {
	var N, M, K, X int
	fmt.Scanln(&N, &M, &K, &X)

	info := make([][]int, N+1)
	arrived := make([]bool, N+1)
	distance := make([]int, N+1)

	for i := 0; i < len(distance); i++ {
		distance[i] = -1
	}

	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		info[a] = append(info[a], b)
	}

	arrived[X] = true
	distance[X] = 0

	queue := list.New()
	queue.PushBack(X)
	for queue.Len() > 0 {
		item := queue.Front()
		now := item.Value.(int)
		queue.Remove(item)

		for _, v := range info[now] {
			if !arrived[v] {
				queue.PushBack(v)
				arrived[v] = true
				distance[v] = distance[now] + 1
			}
		}
	}

	cityExisted := false

	for i := 1; i < len(distance); i++ {
		if distance[i] == K {
			fmt.Println(i)
			cityExisted = true
		}
	}

	if !cityExisted {
		fmt.Println(-1)
	}
}
