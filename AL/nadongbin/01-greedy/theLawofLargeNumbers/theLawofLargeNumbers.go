package main

import (
	"fmt"
	"sort"
)

func main() {
	n, m, k := 5, 8, 4
	nums := []int{7, 6, 5, 1, 2}
	fmt.Println(solution(n, m, k, nums))
}

func solution(n, m, k int, nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	count := (m / (k + 1)) * k
	count += m % (k + 1)

	sum := 0
	sum += count * nums[0]
	sum += (m - count) * nums[1]

	return sum
}
