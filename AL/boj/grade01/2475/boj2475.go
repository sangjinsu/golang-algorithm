package main

import "fmt"

func main() {
	var nums [5]uint

	_, err := fmt.Scanf("%d %d %d %d %d", &nums[0], &nums[1], &nums[2], &nums[3], &nums[4])
	if err != nil {
		return
	}

	var sum uint
	for _, num := range nums {
		sum += num * num
	}
	result := sum % 10

	fmt.Println(result)
}
