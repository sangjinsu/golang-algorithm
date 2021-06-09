package main

import "fmt"

func main() {
	var a, b uint8

	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		return
	}

	fmt.Println(a * b)
}
