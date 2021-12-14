package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	mat := make([][3]int, N)
	for i := 0; i < N; i++ {
		var R, G, B int
		fmt.Fscanf(reader, "%d %d %d\n", &R, &G, &B)
		mat[i][0], mat[i][1], mat[i][2] = R, G, B
	}

	for i := 1; i < N; i++ {
		if mat[i-1][1] > mat[i-1][2] {
			mat[i][0] += mat[i-1][2]
		} else {
			mat[i][0] += mat[i-1][1]
		}

		if mat[i-1][0] > mat[i-1][2] {
			mat[i][1] += mat[i-1][2]
		} else {
			mat[i][1] += mat[i-1][0]
		}

		if mat[i-1][0] > mat[i-1][1] {
			mat[i][2] += mat[i-1][1]
		} else {
			mat[i][2] += mat[i-1][0]
		}
	}

	result := 1000 * 1000
	for _, num := range mat[N-1] {
		if num < result {
			result = num
		}
	}
	fmt.Fprint(writer, result)
}
