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

	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = i
	}

	for i := 1; i < n+1; i++ {
		num := 2
		for i >= num*num {
			if dp[i] > dp[i-num*num]+1 {
				dp[i] = dp[i-num*num] + 1
			}
			num++
		}
	}

	fmt.Fprintln(writer, dp[n])
}
