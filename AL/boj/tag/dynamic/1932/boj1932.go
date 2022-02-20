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
	fmt.Fscanln(reader, &N)

	mat := make([][]int, N)
	for i := 0; i < N; i++ {
		mat[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Fscan(reader, &mat[i][j])
		}
	}

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}

	dp[0][0] = mat[0][0]

	for i := 0; i < N-1; i++ {

		for j := 0; j < i+1; j++ {
			if dp[i+1][j] < mat[i+1][j]+dp[i][j] {
				dp[i+1][j] = mat[i+1][j] + dp[i][j]
			}
			if dp[i+1][j+1] < mat[i+1][j+1]+dp[i][j] {
				dp[i+1][j+1] = mat[i+1][j+1] + dp[i][j]
			}
		}
	}

	result := 0
	for i := 0; i < N; i++ {
		if result < dp[N-1][i] {
			result = dp[N-1][i]
		}
	}

	fmt.Fprintln(writer, result)
}
