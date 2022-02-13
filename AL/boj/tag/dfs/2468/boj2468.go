package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

var dx = []int{0, 1, 0, -1}
var dy = []int{-1, 0, 1, 0}
var N int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	mat := make([][]int, N)
	var max int
	for i := 0; i < N; i++ {
		var num int
		for j := 0; j < N; j++ {
			fmt.Fscan(reader, &num)
			mat[i] = append(mat[i], num)
			if max < num {
				max = num
			}
		}
	}

	var result int

	for rain := 0; rain < max; rain++ {
		temp := make([][]int, N)
		for j := 0; j < N; j++ {
			temp[j] = make([]int, N)
			copy(temp[j], mat[j])
		}

		for y := 0; y < N; y++ {
			for x := 0; x < N; x++ {
				if temp[y][x] <= rain {
					temp[y][x] = -1
				}
			}
		}

		var cnt int
		for y := 0; y < N; y++ {
			for x := 0; x < N; x++ {
				if temp[y][x] > 0 {
					bfs(y, x, temp)
					cnt += 1
				}
			}
		}

		if result < cnt {
			result = cnt
		}

	}
	fmt.Fprintln(writer, result)
}

type crd struct {
	x int
	y int
}

func bfs(y int, x int, temp [][]int) {
	queue := list.New()
	queue.PushBack(crd{x: x, y: y})
	temp[y][x] = -1
	for queue.Len() > 0 {
		e := queue.Front()
		nx, ny := e.Value.(crd).x, e.Value.(crd).y
		queue.Remove(e)
		for i := 0; i < 4; i++ {
			ty, tx := ny+dy[i], nx+dx[i]
			if 0 > ty || ty >= N {
				continue
			}
			if 0 > tx || tx >= N {
				continue
			}
			if temp[ty][tx] <= 0 {
				continue
			}
			temp[ty][tx] = -1
			queue.PushBack(crd{x: tx, y: ty})
		}
	}

}
