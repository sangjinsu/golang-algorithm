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

	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	mat := make([][]string, n)
	for i := 0; i < n; i++ {
		var line string
		fmt.Fscanln(reader, &line)
		wb := strings.Split(line, "")
		mat[i] = wb
	}

	var whiteChecker []string
	var blackChecker []string

	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			whiteChecker = append(whiteChecker, "WBWBWBWB")
			blackChecker = append(blackChecker, "BWBWBWBW")
		} else {
			whiteChecker = append(whiteChecker, "BWBWBWBW")
			blackChecker = append(blackChecker, "WBWBWBWB")
		}
	}

	result := 64
	for i := 0; i < n-8+1; i++ {
		for j := 0; j < m-8+1; j++ {
			var black, white int
			for k := i; k < i+8; k++ {
				for l := j; l < j+8; l++ {
					if mat[k][l] != string(whiteChecker[k-i][l-j]) {
						white += 1
					}

					if mat[k][l] != string(blackChecker[k-i][l-j]) {
						black += 1
					}
				}
			}

			if result > black {
				result = black
			}

			if result > white {
				result = white
			}
		}
	}

	fmt.Fprintf(writer, "%d", result)

}
