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
		dp = append(dp, make([]int, 2))
	}

	dp[1] = []int{0, 1}
	for i := 2; i < n+1; i++ {
		dp[i][0] = dp[i-1][0] + dp[i-1][1]
		dp[i][1] = dp[i-1][0]
	}

	result := dp[n][0] + dp[n][1]

	fmt.Fprint(writer, result)
}
