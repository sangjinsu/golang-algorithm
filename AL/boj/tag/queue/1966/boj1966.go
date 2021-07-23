package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var testCase int
	fmt.Fscanln(reader, &testCase)

	for i := 0; i < testCase; i++ {
		var n, m int
		fmt.Fscanf(reader, "%d %d\n", &n, &m)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		numsStr := strings.Split(input, " ")
		nums := list.New()
		for _, n := range numsStr {
			num, _ := strconv.Atoi(n)
			nums.PushBack(num)
		}

		var count int
		for true {
			max := findMax(nums)

			item := nums.Front()
			nums.Remove(item)

			if item.Value != max {
				nums.PushBack(item.Value.(int))
			} else if m == 0 {
				count++
				break
			} else {
				count++
				n--
				max = findMax(nums)
			}
			m--
			if m < 0 {
				m += n
			}
		}

		fmt.Fprintln(writer, count)

	}
}

func findMax(nums *list.List) int {
	var max int
	for num := nums.Front(); num != nil; num = num.Next() {
		if max < num.Value.(int) {
			max = num.Value.(int)
		}
	}
	return max
}
