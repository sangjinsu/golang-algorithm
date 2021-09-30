package main

import (
	"bufio"
	"fmt"
	"os"
)

var dx = []int{1, 2, 2, 1, -1, -2, -2, -1}
var dy = []int{-2, -1, 1, 2, 2, 1, -1, -2}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var I, x, y, ax, ay int
		fmt.Fscanln(reader, &I)
		fmt.Fscanf(reader, "%d %d\n", &x, &y)
		fmt.Fscanf(reader, "%d %d\n", &ax, &ay)

		mat := make([][]int, I)
		for j := 0; j < I; j++ {
			mat[j] = make([]int, I)
		}

		bfs(y, x, mat, I, ay, ax)
		fmt.Fprintln(writer, mat[ay][ax]-1)
	}

}

func bfs(y int, x int, mat [][]int, I int, ay, ax int) {
	mat[y][x] = 1
	if y == ay && x == ax {
		return
	}

	queue := [][]int{
		{y, x},
	}
	for len(queue) > 0 {
		ny, nx := queue[0][0], queue[0][1]
		queue = queue[1:]
		for i := 0; i < 8; i++ {
			ty, tx := ny+dy[i], nx+dx[i]
			if 0 <= ty && ty < I && 0 <= tx && tx < I && mat[ty][tx] == 0 {
				mat[ty][tx] = mat[ny][nx] + 1
				if ty == ay && tx == ax {
					return
				}
				queue = append(queue, []int{ty, tx})
			}
		}
	}
}
