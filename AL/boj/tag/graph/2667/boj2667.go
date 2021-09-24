package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var dx = []int{0, 1, 0, -1}
var dy = []int{-1, 0, 1, 0}

func dfs(y, x int, mat [][]int, N int) int {
	mat[y][x] = -1
	stack := [][]int{
		{y, x},
	}
	cnt := 1
	for len(stack) > 0 {
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ny, nx := pos[0], pos[1]
		for i := 0; i < 4; i++ {
			ty, tx := ny+dy[i], nx+dx[i]
			if 0 <= ty && ty < N && 0 <= tx && tx < N && mat[ty][tx] == 1 {
				stack = append(stack, []int{ty, tx})
				mat[ty][tx] = -1
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)

	mat := make([][]int, N)
	for i := 0; i < N; i++ {
		var line string
		fmt.Fscanln(reader, &line)
		inputs := strings.Split(line, "")
		temp := make([]int, N)
		for j, input := range inputs {
			num, _ := strconv.Atoi(input)
			temp[j] = num
		}
		mat[i] = temp
	}

	var result []int
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if mat[y][x] == 1 {
				result = append(result, dfs(y, x, mat, N))
			}
		}
	}

	sort.Ints(result)
	fmt.Fprintln(writer, len(result))
	for _, num := range result {
		fmt.Fprintln(writer, num)
	}

}
