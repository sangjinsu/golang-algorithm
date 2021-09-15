package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	var mat [][]string
	for i := 0; i < n; i++ {
		var line string
		fmt.Fscanln(reader, &line)
		mat = append(mat, strings.Split(line, ""))
	}

	result := 0
	var chk int
	for y := 0; y < n; y++ {
		for x := 0; x < n-1; x++ {

			mat[y][x], mat[y][x+1] = mat[y][x+1], mat[y][x]
			chk = check(mat, n)
			if chk > result {
				result = chk
			}
			mat[y][x], mat[y][x+1] = mat[y][x+1], mat[y][x]

			mat[x][y], mat[x+1][y] = mat[x+1][y], mat[x][y]
			chk = check(mat, n)
			if chk > result {
				result = chk
			}
			mat[x][y], mat[x+1][y] = mat[x+1][y], mat[x][y]
		}
	}
	fmt.Fprint(writer, result)
}

func check(mat [][]string, n int) int {
	result := 0
	for y := 0; y < n; y++ {
		row, col := 1, 1
		for x := 1; x < n; x++ {
			if mat[y][x] == mat[y][x-1] {
				row += 1
			} else {
				if row > result {
					result = row
				}
				row = 1
			}

			if mat[x][y] == mat[x-1][y] {
				col += 1
			} else {
				if col > result {
					result = col
				}
				col = 1
			}
		}

		if row > result {
			result = row
		}
		if col > result {
			result = col
		}
	}
	return result
}
