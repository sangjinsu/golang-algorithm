package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	// 상 우 하 좌
	N   int
	dy  = []int{-1, 0, 1, 0}
	dx  = []int{0, 1, 0, -1}
	mat []string
)

func dfs(ny, nx int, visited [][]bool, colors string) {
	for i := 0; i < 4; i++ {
		ty, tx := ny+dy[i], nx+dx[i]
		if 0 > ty || ty >= N {
			continue
		}
		if 0 > tx || tx >= N {
			continue
		}
		if visited[ty][tx] == true {
			continue
		}
		if strings.Contains(colors, string(mat[ty][tx])) {
			visited[ty][tx] = true
			dfs(ty, tx, visited, colors)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d\n", &N)

	mat = make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Fscanf(reader, "%s\n", &mat[i])
	}

	visited := make([][]bool, N)
	for i := 0; i < N; i++ {
		visited[i] = make([]bool, N)
	}

	var cnt int
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if visited[y][x] == false {
				cnt++
				dfs(y, x, visited, string(mat[y][x]))
			}
		}
	}

	fmt.Fprintln(writer, cnt)

	visited = make([][]bool, N)
	for i := 0; i < N; i++ {
		visited[i] = make([]bool, N)
	}
	cnt = 0
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if visited[y][x] == false {
				cnt++
				if strings.Contains("RG", string(mat[y][x])) {
					dfs(y, x, visited, "RG")
				} else {
					dfs(y, x, visited, string(mat[y][x]))
				}
			}
		}
	}

	fmt.Fprintln(writer, cnt)
}
