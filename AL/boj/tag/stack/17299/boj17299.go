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

	var N int
	fmt.Fscanln(reader, &N)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	inputs := strings.Split(input, " ")

	numbers := make([]int, len(inputs))
	for i, n := range inputs {
		num, _ := strconv.Atoi(n)
		numbers[i] = num
	}

	fmap := make(map[int]int)
	for _, number := range numbers {
		fmap[number]++
	}

	var stack []int
	var result []string
	for i := 0; i < len(numbers); i++ {
		result = append(result, "-1")
	}

	for i, _ := range numbers {
		for len(stack) > 0 && fmap[numbers[stack[len(stack)-1]]] < fmap[numbers[i]] {
			result[stack[len(stack)-1]] = strconv.Itoa(numbers[i])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	fmt.Fprint(writer, strings.Join(result, " "))
}
