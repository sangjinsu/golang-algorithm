package main

import (
	"fmt"
)

func main() {
	var x, y int

	_, err := fmt.Scanln(&x)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&y)
	if err != nil {
		return
	}

	quadrant := 0

	if x > 0 {
		if y > 0 {
			quadrant = 1
		} else {
			quadrant = 4
		}
	} else {
		if y > 0 {
			quadrant = 2
		} else {
			quadrant = 3
		}
	}

	fmt.Println(quadrant)
}
