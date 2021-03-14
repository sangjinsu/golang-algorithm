package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)
	strArr := strings.Split(input, " ")

	lengths := make([]int, n)
	for i, length := range strArr {
		lengths[i], _ = strconv.Atoi(length)
	}

	sort.Ints(lengths)

	start := 0
	end := lengths[len(lengths)-1]

	result := 0
	for start <= end {
		mid := int((start + end) / 2)

		sum := 0
		for _, length := range lengths {
			if mid < length {
				sum += length - mid
			}
		}

		if sum < m {
			end = mid - 1
		} else {
			result = mid
			start = mid + 1
		}
	}

	fmt.Println(result)
}
