package main

import "fmt"

func main() {
	var a, b byte

	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		return
	}

	_, err = fmt.Scanf("%d", &b)
	if err != nil {
		return
	}

	fmt.Println(a + b)
}
