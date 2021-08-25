package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var pointers [4][4]int
	for i := 0; i < 4; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		inputs := strings.Split(line, " ")

		var pointer [4]int
		for i, input := range inputs {
			n, _ := strconv.Atoi(input)
			pointer[i] = n
		}

		pointers[i] = pointer
	}

	var plane [101][101]int
	for _, pointer := range pointers {
		x1, y1, x2, y2 := pointer[0], pointer[1], pointer[2], pointer[3]
		for y := y1; y < y2; y++ {
			for x := x1; x < x2; x++ {
				plane[y][x] = 1
			}
		}
	}

	var result int
	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			result += plane[i][j]
		}
	}

	fmt.Fprint(writer, result)
}
