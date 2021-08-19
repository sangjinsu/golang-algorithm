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

	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	n, _ := strconv.Atoi(input)

	dp := make([]int, n+1)
	for i := 2; i < n+1; i++ {
		dp[i] = 1000001
	}

	for i := 2; i < n+1; i++ {
		if i%2 == 0 {
			if dp[i] > dp[i/2]+1 {
				dp[i] = dp[i/2] + 1
			}
		}

		if i%3 == 0 {
			if dp[i] > dp[i/3]+1 {
				dp[i] = dp[i/3] + 1
			}
		}

		if dp[i] > dp[i-1]+1 {
			dp[i] = dp[i-1] + 1
		}
	}

	fmt.Fprintf(writer, "%d", dp[len(dp) - 1])
}
