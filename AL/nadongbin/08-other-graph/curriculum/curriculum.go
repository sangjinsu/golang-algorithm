package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scanln(&N)

	indegree := make([]int, N+1)
	graph := make([][]int, N+1)
	time := make([]int, N+1)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 1; i < N+1; i++ {
		scanner.Scan()
		text := scanner.Text()
		inputStr := strings.Split(text, " ")
		data := make([]int, len(inputStr))
		for j, v := range inputStr {
			data[j], _ = strconv.Atoi(v)
		}
		time[i] = data[0]
		for _, v := range data[1 : len(data)-1] {
			indegree[i]++
			graph[v] = append(graph[v], i)
		}
	}

	topologySort(time, N, indegree, graph)
}

func topologySort(time []int, N int, indegree []int, graph [][]int) {
	result := make([]int, len(time))
	copy(result, time)
	queue := list.New()

	for i := 1; i < N+1; i++ {
		if indegree[i] == 0 {
			queue.PushBack(i)
		}
	}

	fmt.Println(result, time, indegree, graph)

	for queue.Len() > 0 {
		item := queue.Front()
		now := item.Value.(int)
		queue.Remove(item)

		for _, i := range graph[now] {
			result[i] = int(math.Max(float64(result[i]), float64(result[now]+time[i])))
			indegree[i]--
			if indegree[i] == 0 {
				queue.PushBack(i)
			}
		}
	}
	fmt.Println(result, time, indegree, graph)
	for i := 1; i < N+1; i++ {
		fmt.Println(result[i])
	}
}
