package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type Pos struct {
	x int
	y int
}

var (
	N      int
	M      int
	mat    [][]int
	result int
	dx     = []int{0, 0, -1, 1}
	dy     = []int{-1, 1, 0, 0}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	mat = make([][]int, N)
	for y := 0; y < N; y++ {
		line := make([]int, M)
		for x := 0; x < M; x++ {
			fmt.Fscan(reader, &line[x])
		}
		mat[y] = line
	}

	recursion(0)

	fmt.Fprintf(writer, "%d\n", result)
}

func recursion(block int) {
	if block == 3 {
		cnt := bfs(mat)
		if result < cnt {
			result = cnt
		}
		return
	}

	for y := 0; y < N; y++ {
		for x := 0; x < M; x++ {
			if mat[y][x] == 0 {
				mat[y][x] = 1
				recursion(block + 1)
				mat[y][x] = 0
			}
		}
	}
}

func bfs(mat [][]int) int {
	queue := list.New()
	m := make([][]int, N)

	for y := 0; y < N; y++ {
		line := make([]int, M)
		copy(line, mat[y])
		m[y] = line
		for x := 0; x < M; x++ {
			if m[y][x] == 2 {
				queue.PushBack(Pos{x, y})
			}
		}
	}

	for queue.Len() > 0 {
		item := queue.Front()
		pos := item.Value.(Pos)
		queue.Remove(item)

		ny, nx := pos.y, pos.x
		for i := 0; i < 4; i++ {
			ty, tx := ny+dy[i], nx+dx[i]
			if ty < 0 || ty >= N {
				continue
			}
			if tx < 0 || tx >= M {
				continue
			}
			if m[ty][tx] > 0 {
				continue
			}

			m[ty][tx] = 2
			queue.PushBack(Pos{tx, ty})
		}
	}

	var cnt int

	for y := 0; y < N; y++ {
		for x := 0; x < M; x++ {
			if m[y][x] == 0 {
				cnt++
			}
		}
	}

	return cnt
}
