package main

import "fmt"

func main() {
	var a, b float64
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		return
	}
	fmt.Println(a/b)
}
