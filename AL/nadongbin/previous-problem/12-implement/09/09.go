package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	s := "aabbaccc"
	solution(s)
}

func solution(input string) {
	answer := len(input)
	for step := 1; step < len(input)/2+1; step++ {
		com := ""
		prev := input[0:step]
		count := 1

		for i := step; i < len(input); i += step {

			nextStep := i + step
			if nextStep > len(input) {
				nextStep = len(input)
			}

			if prev == input[i:nextStep] {
				count++
			} else {
				if count >= 2 {
					com += strconv.Itoa(count) + prev
				} else {
					com += prev
				}

				prev = input[i:nextStep]
				count = 1
			}
		}

		if count >= 2 {
			com += strconv.Itoa(count) + prev
		} else {
			com += prev
		}

		answer = int(math.Min(float64(len(com)), float64(answer)))
	}

	fmt.Println(answer)
}
