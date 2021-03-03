package main

import "fmt"

func main() {
	var n int = 1260
	var cnt int

	coinTypes := []int{500, 100, 50, 10}

	for _, coin := range coinTypes {
		cnt += n / coin
		n %= coin
	}
	fmt.Println(cnt)
}
