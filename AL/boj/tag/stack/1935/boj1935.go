package main

import (
	"bufio"
	"fmt"
	"os"
)

func solvePostfix(expr string, mapper map[rune]float64) float64 {
	var stack []float64
	for _, e := range expr {
		if 'A' <= e && e <= 'Z' {
			stack = append(stack, mapper[e])
		} else {
			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if e == '+' {
				stack = append(stack, num1+num2)
			} else if e == '-' {
				stack = append(stack, num1-num2)
			} else if e == '*' {
				stack = append(stack, num1*num2)
			} else if e == '/' {
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var expr string
	fmt.Fscanln(reader, &expr)

	mapper := make(map[rune]float64)
	var num float64
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &num)
		mapper[rune('A'+i)] = num
	}
	fmt.Println(mapper)
	result := solvePostfix(expr, mapper)
	fmt.Fprint(writer, result)
}
