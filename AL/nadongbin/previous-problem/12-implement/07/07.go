package main

import (
	"fmt"
	"strconv"
)

func main() {
	inputs := []int{123402, 7755}

	for i := 0; i < len(inputs); i++ {
		solution(inputs[i])
	}

}

func sum(num string, c chan int) {
	sum := 0
	for i := 0; i < len(num); i++ {
		n, _ := strconv.Atoi(string(num[i]))
		sum += n
	}
	c <- sum
}

func solution(input int) {
	numStr := strconv.Itoa(input)
	left := numStr[:len(numStr)/2]
	right := numStr[len(numStr)/2:]

	c := make(chan int, 2)

	go sum(left, c)
	go sum(right, c)

	leftSum := <-c
	rightSum := <-c

	if leftSum == rightSum {
		fmt.Println("LUCKY")
	} else {
		fmt.Println("READY")
	}
}
