package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

var dx = []int{0, 1, 0, -1}
var dy = []int{-1, 0, 1, 0}

func bfs(queue *list.List, mat [][]int, M, N int) {
	for queue.Len() > 0 {
		node := queue.Front()
		ny, nx := node.Value.(pos).y, node.Value.(pos).x
		queue.Remove(node)
		for i := 0; i < 4; i++ {
			ty, tx := ny+dy[i], nx+dx[i]
			if 0 <= ty && ty < N && 0 <= tx && tx < M && mat[ty][tx] == 0 {
				mat[ty][tx] = mat[ny][nx] + 1
				queue.PushBack(pos{x: tx, y: ty})
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var M, N int
	fmt.Fscanf(reader, "%d %d\n", &M, &N)

	mat := make([][]int, N)
	for i := 0; i < N; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		inputs := strings.Split(line, " ")

		nums := make([]int, M)
		for j, input := range inputs {
			num, _ := strconv.Atoi(input)
			nums[j] = num
		}
		mat[i] = nums
	}

	queue := list.New()
	for y := 0; y < N; y++ {
		for x := 0; x < M; x++ {
			if mat[y][x] == 1 {
				queue.PushBack(pos{x: x, y: y})
			}
		}
	}
	bfs(queue, mat, M, N)
	result := 0
	for _, line := range mat {
		for _, v := range line {
			if v == 0 {
				result = -1
				break
			}
			if result < v {
				result = v
			}
		}
		if result < 0 {
			break
		}
	}
	if result != -1 {
		result--
	}

	fmt.Fprintln(writer, result)
}
