package main

import (
	"fmt"
	"math"
)

func main() {
	var input []byte
	fmt.Scanln(&input)

	count0 := 0
	count1 := 0

	if string(input[0]) == "0" {
		count1++
	} else {
		count0++
	}

	for i := 0; i < len(input)-1; i++ {
		if input[i] != input[i+1] {
			if string(input[i+1]) == "0" {
				count1++
			} else {
				count0++
			}
		}
	}

	fmt.Println(int(math.Min(float64(count0), float64(count1))))
}
