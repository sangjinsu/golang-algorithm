package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1_000_000_009
const max = 1_000_001

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var dp [max]int
	dp[1], dp[2], dp[3] = 1, 2, 4
	for i := 4; i < max; i++ {
		dp[i] = (dp[i-3] + dp[i-2] + dp[i-1])%mod
	}

	var t int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		fmt.Fprint(writer, strconv.Itoa(dp[n]) + "\n")
	}
}
