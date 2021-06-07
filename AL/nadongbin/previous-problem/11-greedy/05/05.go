package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, M int
	fmt.Scanln(&N, &M)

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()
	inputStr := strings.Split(input, " ")
	nums := make([]int, len(inputStr))
	for i, v := range inputStr {
		nums[i], _ = strconv.Atoi(v)
	}

	// count := 0
	// for i := 0; i < len(nums); i++ {
	// 	for j := i; j < len(nums); j++ {
	// 		if nums[i] != nums[j] {
	// 			count++
	// 		}
	// 	}
	// }

	arr := make([]int, 11)
	for _, v := range nums {
		arr[v]++
	}

	count := 0

	for i := 1; i < M+1; i++ {
		N -= arr[i]
		count += arr[i] * N
	}

	fmt.Println(count)
}
