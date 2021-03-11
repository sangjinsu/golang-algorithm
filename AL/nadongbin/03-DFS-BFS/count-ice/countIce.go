package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countIce(graph [][]int, x, y, n, m int) bool {
	if x <= -1 || x >= n || y <= -1 || y >= m {
		return false
	}
	if graph[x][y] == 0 {
		graph[x][y] = 1
		countIce(graph, x-1, y, n, m)
		countIce(graph, x+1, y, n, m)
		countIce(graph, x, y-1, n, m)
		countIce(graph, x, y+1, n, m)
		return true
	}
	return false
}

func main() {
	result := 0
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

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if countIce(graph, i, j, n, m) == true {
				result++
			}
		}
	}
	fmt.Println(result)
}
