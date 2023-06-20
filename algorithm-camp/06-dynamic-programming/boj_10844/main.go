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

	dp := make([][]uint, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]uint, 10)
	}

	for i := 1; i < 10; i++ {
		dp[1][i] = 1
	}

	for i := 2; i < n+1; i++ {
		for j := 0; j < 10; j++ {
			switch j {
			case 0:
				dp[i][j] += dp[i-1][j+1]

			case 9:
				dp[i][j] += dp[i-1][j-1]
			default:
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j+1]
			}
			dp[i][j] %= 1_000_000_000
		}
	}

	var result uint
	for _, num := range dp[n] {
		result += num
	}
	result %= 1_000_000_000

	fmt.Fprint(writer, result)
}
