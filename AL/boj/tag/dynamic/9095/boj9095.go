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

	var t int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)

		dp := make([]int, 11)
		dp[1], dp[2], dp[3] = 1, 2, 4
		for i := 4; i < 11; i++ {
			dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
		}

		fmt.Fprintln(writer, dp[n])
	}
}
