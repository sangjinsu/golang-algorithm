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

	var N, K int
	fmt.Fscanf(reader, "%d %d\n", &N, &K)

	items := make([][]int, N)
	for i := 0; i < N; i++ {
		var w, v int
		fmt.Fscanf(reader, "%d %d\n", &w, &v)
		items[i] = []int{w, v}
	}

	dp := make([]int, K+1)
	for i := 0; i < N; i++ {
		for j := K; j >= 0; j-- {
			if items[i][0] <= j {
				if dp[j] < dp[j-items[i][0]]+items[i][1] {
					dp[j] = dp[j-items[i][0]] + items[i][1]
				}
			}
		}
	}
	fmt.Fprint(writer, dp[K])
}
