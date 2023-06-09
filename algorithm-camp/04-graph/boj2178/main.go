package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordination struct {
	x int
	y int
}

var (
	N int
	M int
	// 상 우 하 좌
	dy = []int{-1, 0, 1, 0}
	dx = []int{0, 1, 0, -1}
)

func bfs(mat [][]int, visited [][]int, y, x int) {
	visited[y][x] = 1
	queue := list.New()
	queue.PushBack(coordination{x: x, y: y})

	for queue.Len() > 0 {
		coord := queue.Remove(queue.Front()).(coordination)
		ny, nx := coord.y, coord.x
		for i := 0; i < 4; i++ {
			ty, tx := ny+dy[i], nx+dx[i]
			if 0 <= ty && ty < N && 0 <= tx && tx < M && mat[ty][tx] == 1 && visited[ty][tx] == 0 {
				visited[ty][tx] = visited[ny][nx] + 1
				queue.PushBack(coordination{x: tx, y: ty})
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	mat := make([][]int, N)
	for i := 0; i < N; i++ {
		var input string
		fmt.Fscanf(reader, "%s\n", &input)
		row := strings.Split(input, "")
		for _, value := range row {
			num, _ := strconv.Atoi(value)
			mat[i] = append(mat[i], num)
		}
	}

	visited := make([][]int, N)
	for i := 0; i < N; i++ {
		visited[i] = append(visited[i], make([]int, M)...)
	}

	bfs(mat, visited, 0, 0)

	fmt.Fprintf(writer, "%d", visited[N-1][M-1])

}
