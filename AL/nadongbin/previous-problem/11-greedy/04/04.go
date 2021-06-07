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
	var N int
	fmt.Scanln(&N)

	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadString('\n')
	input = strings.TrimSuffix(input, "\r\n")

	inputStr := strings.Split(input, " ")

	nums := make([]int, len(inputStr))
	for i, s := range inputStr {
		nums[i], _ = strconv.Atoi(s)
	}

	// sort.Slice(nums, func(i, j int) bool {
	// 	return nums[i] > nums[j]
	// })

	// result := 1
	// for {
	// 	pay := result
	// 	for _, num := range nums {
	// 		if pay >= num {
	// 			pay -= num
	// 		}
	// 	}
	// 	if pay != 0 {
	// 		break
	// 	}
	// 	result++
	// }

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result := 1
	for _, num := range nums {
		if result < num {
			break
		}
		result += num
	}

	fmt.Println(result)
}
