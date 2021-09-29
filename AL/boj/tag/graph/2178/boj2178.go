package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dx = []int{0, 1, 0, -1}
var dy = []int{-1, 0, 1, 0}

func bfs(y, x int, mat [][]int, N, M int) {
	mat[y][x] = 2
	queue := list.New()
	queue.PushBack([]int{y, x})

	for queue.Len() > 0 {
		pos := queue.Front()
		ny, nx := pos.Value.([]int)[0], pos.Value.([]int)[1]
		queue.Remove(pos)

		for i := 0; i < 4; i++ {
			ty, tx := ny+dx[i], nx+dy[i]
			if 0 <= ty && ty < N && 0 <= tx && tx < M && mat[ty][tx] == 1 {
				mat[ty][tx] = mat[ny][nx] + 1
				queue.PushBack([]int{ty, tx})
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	mat := make([][]int, N)
	for i := 0; i < N; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		inputs := strings.Split(line, "")
		nums := make([]int, M)
		for j, input := range inputs {
			num, _ := strconv.Atoi(input)
			nums[j] = num
		}
		mat[i] = nums
	}

	bfs(0, 0, mat, N, M)
	fmt.Fprint(writer, mat[N-1][M-1]-1)
}
