package main

import (
	"bufio"
	"fmt"
	"os"
)

var cache = make(map[int]int)

func recursion(n int) int {
	if v, ok := cache[n]; ok {
		return v
	}

	if n <= 1 {
		cache[n] = 1
		return 1
	}

	cache[n] = (recursion(n-1) + recursion(n-2)) % 10_007
	return cache[n]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	dp := make([]uint, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i < n+1; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 10_007
	}

	recursion(n)

	fmt.Fprint(writer, cache[n])
}
