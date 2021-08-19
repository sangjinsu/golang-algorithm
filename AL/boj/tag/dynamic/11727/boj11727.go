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

	dp := make([]uint64, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		dp[i] = (dp[i-2]*2 + dp[i-1]) % 10007
	}
	fmt.Fprintln(writer, dp[n])
}
