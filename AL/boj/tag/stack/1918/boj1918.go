package main

import (
	"bufio"
	"fmt"
	"os"
)

func makePostfix(expr string) string {
	op := map[rune]int{
		'(': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	var stack []rune
	var postfix string
	for _, e := range expr {
		if 'A' <= e && e <= 'Z' {
			postfix += string(e)
		} else {
			if e == '(' {
				stack = append(stack, e)
			} else if e == ')' {
				for len(stack) > 0 && stack[len(stack)-1] != '(' {
					postfix += string(stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = stack[:len(stack)-1]
			} else {
				for len(stack) > 0 && op[e] <= op[stack[len(stack)-1]] {
					postfix += string(stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, e)
			}
		}
	}

	for len(stack) > 0 {
		postfix += string(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var expr string
	fmt.Fscanln(reader, &expr)

	result := makePostfix(expr)
	fmt.Fprint(writer, result)
}
