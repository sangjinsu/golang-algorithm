package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type virus struct {
	x   int
	y   int
	num int
}

func main() {
	var n, k int
	fmt.Scanln(&n, &k)

	graph := make([][]int, n)
	virusInfo := []virus{}

	for i := 0; i < n; i++ {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		inputStr := s.Text()
		input := strings.Split(inputStr, " ")
		for j := 0; j < len(input); j++ {
			item, _ := strconv.Atoi(input[j])
			graph[i] = append(graph[i], item)
			if item != 0 {
				virusInfo = append(virusInfo, virus{x: i, y: j, num: item})
			}
		}
	}

	sort.Slice(virusInfo, func(i, j int) bool {
		return virusInfo[i].num < virusInfo[j].num
	})

	queue := list.New()
	for i := 0; i < len(virusInfo); i++ {
		queue.PushBack(virusInfo[i])
	}

	var s, x, y int
	fmt.Scanln(&s, &x, &y)

	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	for queue.Len() > 0 {
		if s == 0 {
			break
		}

		item := queue.Front()
		now := item.Value.(virus)
		queue.Remove(item)
		for i := 0; i < 4; i++ {
			nx := now.x + dx[i]
			ny := now.y + dy[i]

			if nx > 0 && nx < n && ny > 0 && ny < n && graph[nx][ny] == 0 {
				graph[nx][ny] = now.num
				queue.PushBack(virus{x: nx, y: ny, num: now.num})
			}
		}

		s--
	}

	result := graph[x-1][y-1]
	fmt.Println(result)
}
