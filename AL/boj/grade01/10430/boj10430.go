package main

import "fmt"

func main() {
	var a, b, c int

	_, err := fmt.Scanf("%d %d %d", &a, &b, &c)
	if err != nil {
		return
	}

	fmt.Println((a + b) % c)
	fmt.Println(((a % c) + (b % c)) % c)
	fmt.Println((a * b) % c)
	fmt.Println(((a % c) * (b % c)) % c)
}
