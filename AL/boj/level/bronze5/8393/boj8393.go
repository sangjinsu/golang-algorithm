package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}

	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println(sum)
}
