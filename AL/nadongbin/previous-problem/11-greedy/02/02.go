package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)

	inputStr := strings.Split(s, "")
	input := make([]int, len(inputStr))
	for i, v := range inputStr {
		input[i], _ = strconv.Atoi(v)
	}

	result := input[0]

	for i := 1; i < len(input); i++ {
		result = int(math.Max(float64(result*input[i]), float64(result+input[i])))
	}

	fmt.Println(result)
}
