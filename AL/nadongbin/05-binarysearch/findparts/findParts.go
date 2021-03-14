package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	partNumbers := make([]int, n)

	var input string
	var inputArr []string
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	input = strings.TrimSpace(input)
	inputArr = strings.Split(input, " ")

	for i, num := range inputArr {
		partNumbers[i], err = strconv.Atoi(num)
		if err != nil {
			log.Fatal("Can not convert string to int")
		}
	}

	var m int
	fmt.Scanln(&m)

	checkNumbers := make([]int, m)

	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	input = strings.TrimSpace(input)
	inputArr = strings.Split(input, " ")

	for i, num := range inputArr {
		checkNumbers[i], err = strconv.Atoi(num)
		if err != nil {
			log.Fatal("Can not convert string to int")
		}
	}

	sort.Slice(partNumbers, func(i, j int) bool { return partNumbers[i] < partNumbers[j] })
	for _, num := range checkNumbers {
		if found := sort.SearchInts(partNumbers, num); found < len(partNumbers) && partNumbers[found] == num {
			fmt.Printf("yes ")
		} else {
			fmt.Printf("no ")
		}
	}
}
