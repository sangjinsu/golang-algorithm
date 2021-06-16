package main

import (
	"fmt"
)

func main() {
	var input string

	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return
	}

	ascii := int(input[0])

	fmt.Println(ascii)
}
