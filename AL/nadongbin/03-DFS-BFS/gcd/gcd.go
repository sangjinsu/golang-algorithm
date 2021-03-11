package main

import "fmt"

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func main() {
	fmt.Println(gcd(162, 192))
}
