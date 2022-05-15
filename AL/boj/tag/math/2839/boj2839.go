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
	fmt.Fscanf(reader, "%d", &N)

	dp := make([]int, 5001)
	dp[3], dp[5] = 1, 1

	for i := 6; i < N+1; i++ {
		if i%5 == 0 {
			dp[i] = dp[i-5] + 1
		} else if i%3 == 0 {
			dp[i] = dp[i-3] + 1
		} else if dp[i-3] > 0 && dp[i-5] > 0 {
			if dp[i-3] > dp[i-5] {
				dp[i] = dp[i-3] + 1
			} else {
				dp[i] = dp[i-5] + 1
			}
		}
	}

	if dp[N] == 0 {
		fmt.Fprint(writer, -1)
		return
	}
	fmt.Fprint(writer, dp[N])
}
