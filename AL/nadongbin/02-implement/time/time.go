package main

import (
	"fmt"
)

func main() {
	input := 23

	count := 0

	for hour := 0; hour <= input; hour++ {
		for minute := 0; minute < 60; minute++ {
			for secound := 0; secound < 60; secound++ {
				if hour%10 == 3 || minute%10 == 3 || minute/10 == 3 || secound%10 == 3 || secound/10 == 3 {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
