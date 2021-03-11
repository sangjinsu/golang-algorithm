package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type cdn struct {
	x int
	y int
}

func escape(graph [][]int, n, m int) int {
	queue := list.New()
	queue.PushBack(cdn{x: 0, y: 0})

	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}

	for queue.Len() > 0 {
		e := queue.Front()
		x, y := e.Value.(cdn).x, e.Value.(cdn).y
		queue.Remove(e)

		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]

			if nx <= -1 || nx >= n || ny <= -1 || ny >= m {
				continue
			}

			if graph[nx][ny] == 0 {
				continue
			}

			if graph[nx][ny] == 1 {
				graph[nx][ny] = graph[x][y] + 1
				queue.PushBack(cdn{x: nx, y: ny})
			}
		}
	}

	return graph[n-1][m-1]
}

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	var graph [][]int

	for i := 0; i < n; i++ {
		var str string
		fmt.Scanln(&str)
		strRow := strings.Split(str, "")
		row := make([]int, len(strRow))
		for i, v := range strRow {
			row[i], _ = strconv.Atoi(v)
		}
		graph = append(graph, row)
	}
	fmt.Println(graph)

	fmt.Println(escape(graph, n, m))
}
