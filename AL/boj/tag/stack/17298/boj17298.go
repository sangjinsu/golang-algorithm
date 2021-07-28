package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var k int
	fmt.Fscanln(reader, &k)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	inputNums := strings.Split(input, " ")

	var nums []int
	for _, num := range inputNums {
		n, _ := strconv.Atoi(num)
		nums = append(nums, n)
	}

	stack := []int{0}
	var result []string
	for i := 0; i < k; i++ {
		result = append(result, "-1")
	}

	for i := 1; i < k; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			result[stack[len(stack)-1]] = strconv.Itoa(nums[i])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	fmt.Fprint(writer, strings.Join(result, " "))
}
