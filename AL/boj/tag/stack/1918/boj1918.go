package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isOperatorOrBracket(expr string) bool {
	return expr != "+" && expr != "-" && expr != "*" && expr != "/" && expr != "(" && expr != ")"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input, _ := reader.ReadString('\n')
	expression := strings.Split(strings.TrimSuffix(input, "\n"), "")

	op := map[string]int{
		"(": 0,
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	var stack []string
	var postfix string
	for _, expr := range expression {
		if isOperatorOrBracket(expr) {
			if expr == ")" {
				for len(stack) > 0 && stack[len(stack)-1] != "(" {
					postfix += stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
				stack = stack[:len(stack)-1]
			} else {
				for len(stack) > 0 && op[stack[len(stack)-1]] >= op[expr] {
					postfix += stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, expr)
			}
		} else {
			postfix += expr
		}
	}

	for len(stack) > 0 {
		postfix += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	fmt.Fprint(writer, postfix)
}
