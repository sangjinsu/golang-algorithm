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

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	matrix := make([][]int, N)
	var num int
	for i := 0; i < N; i++ {
		matrix[i] = make([]int, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Fscan(reader, &num)
			matrix[i][j] += num
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Fscan(reader, &num)
			matrix[i][j] += num
		}
	}

	for _, ints := range matrix {
		nums := make([]string, 0, len(ints))
		for _, num := range ints {
			itoa := strconv.Itoa(num)
			nums = append(nums, itoa)
		}
		fmt.Fprintln(writer, strings.Join(nums, " "))
	}
}
