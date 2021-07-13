package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	reader.Scan()
	testCases, _ := strconv.Atoi(reader.Text())

	for i := 0; i < testCases; i++ {
		reader.Scan()
		input := reader.Text()

		stack := list.New()
		check := "YES"
		for _, v := range input {
			letter := string(v)
			if letter == "(" {
				stack.PushBack(letter)
			} else {
				if stack.Len() > 0 {
					item := stack.Back()
					stack.Remove(item)
				} else {
					check = "NO"
					break
				}
			}
		}

		if stack.Len() > 0 {
			check = "NO"
		}

		fmt.Fprintln(writer, check)
	}
}
