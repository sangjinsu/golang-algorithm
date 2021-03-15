package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	money := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&money[i])
	}

	dp := make([]int, m+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = 10001
	}

	for i := 0; i < len(dp); i++ {
		for _, money := range money {
			if i >= money {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-money]+1)))
			}
		}
		fmt.Println(dp)
	}

	if dp[len(dp)-1] <= 10000 {
		fmt.Println(dp[len(dp)-1])
	} else {
		fmt.Println(-1)
	}
}
