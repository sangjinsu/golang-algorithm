package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dx = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dy = []int{-1, 0, 1, 0, -1, 1, 1, -1}

func dfs(ny, nx int, mat [][]int, w, h int) {
	mat[ny][nx] = 2
	for i := 0; i < 8; i++ {
		ty, tx := ny+dy[i], nx+dx[i]
		if 0 <= ty && ty < h && 0 <= tx && tx < w && mat[ty][tx] == 1 {
			dfs(ty, tx, mat, w, h)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for true {
		var w, h int
		fmt.Fscanf(reader, "%d %d\n", &w, &h)

		if w == 0 && h == 0 {
			break
		}

		mat := make([][]int, h)
		for i := 0; i < h; i++ {
			line, _ := reader.ReadString('\n')
			inputs := strings.Split(strings.TrimSuffix(line, "\n"), " ")

			nums := make([]int, w)
			for j, input := range inputs {
				num, _ := strconv.Atoi(input)
				nums[j] = num
			}

			mat[i] = nums
		}

		var cnt int
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				if mat[y][x] == 1 {
					cnt++
					dfs(y, x, mat, w, h)
				}
			}
		}

		fmt.Fprintln(writer, cnt)
	}
}
