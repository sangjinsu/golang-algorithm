package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	fmt.Scanln(&x)

	dp := make([]float64, 30001)

	for i := 2; i < x+1; i++ {
		dp[i] = dp[i-1] + 1
		if i%2 == 0 {
			dp[i] = math.Min(dp[i], dp[i/2]+1)
		}
		if i%3 == 0 {
			dp[i] = math.Min(dp[i], dp[i/3]+1)
		}
		if i%5 == 0 {
			dp[i] = math.Min(dp[i], dp[i/5]+1)
		}
		fmt.Println(dp[:x+1])
	}

	fmt.Println(dp[x])

}
