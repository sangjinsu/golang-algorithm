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
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	nums := make([]int, N)

	line, _ := reader.ReadString('\n')
	values := strings.Fields(line)
	for j, value := range values {
		num, _ := strconv.Atoi(value)
		nums[j] = num
	}

	uniqueNums := make(map[int]int)
	for _, num := range nums {
		if _, ok := uniqueNums[num]; !ok {
			uniqueNums[num] = num
		}
	}

	uniques := make([]int, 0, len(nums))
	for _, num := range uniqueNums {
		uniques = append(uniques, num)
	}

	sortedNums := make([]int, len(uniques))
	copy(sortedNums, uniques)
	sort.Ints(sortedNums)

	indices := make(map[int]string)
	for i, num := range sortedNums {
		if _, ok := indices[num]; !ok {
			indices[num] = strconv.Itoa(i)
		}
	}

	answer := make([]string, N)
	for i, num := range nums {
		answer[i] = indices[num]
	}

	fmt.Fprintln(writer, strings.Join(answer, " "))
}
