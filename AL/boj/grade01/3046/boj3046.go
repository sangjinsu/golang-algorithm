package main

import "fmt"

func main() {
	var r1, r2, s int

	_, err := fmt.Scanf("%d %d", &r1, &s)
	if err != nil {
		return
	}

	r2 = s*2 - r1

	fmt.Println(r2)
}
