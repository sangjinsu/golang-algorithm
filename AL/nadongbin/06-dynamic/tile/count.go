package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	dp := make([]int, n+1)

	dp[1] = 1
	dp[2] = 3

	for i := 3; i < n+1; i++ {
		dp[i] = (dp[i-2]*2 + dp[i-1]) % 796796
	}

	fmt.Println(dp[n])
}
