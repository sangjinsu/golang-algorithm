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

	var n int
	fmt.Fscanln(reader, &n)

	var dp [][]int
	for i := 0; i < n+1; i++ {
		dp = append(dp, make([]int, 10))
	}
	dp[1] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	mod := 1_000_000_000
	for i := 2; i < n+1; i++ {
		for j := 0; j < 10; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j+1]
			} else if 1 <= j && j <= 8 {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j+1]%mod
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	var result int
	for i := 1; i < 10; i++ {
		result += dp[n][i]
	}

	fmt.Fprintln(writer, result%mod)
}
