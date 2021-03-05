package main

import "fmt"

func main() {
	N, K := 17, 3
	fmt.Println(untilValueReachesOne(N, K))
}

func untilValueReachesOne(N int, K int) int {
	count := 0
	for N >= K {
		count += N % K
		N = N / K
		count++
	}

	count += N - 1

	return count
}
