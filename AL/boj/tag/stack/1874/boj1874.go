package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			return
		}
	}(writer)

	var n int
	_, err := fmt.Fscanln(reader, &n)
	if err != nil {
		return
	}

	stack := list.New()
	m := 0
	var result []string

	for i := 0; i < n; i++ {
		var num int
		_, err := fmt.Fscanln(reader, &num)
		if err != nil {
			return
		}
		for num > m {
			m++
			stack.PushBack(m)
			result = append(result, "+")
		}

		item := stack.Back()
		if item.Value.(int) == num {
			stack.Remove(item)
			result = append(result, "-")

		} else {
			_, err2 := fmt.Fprintln(writer, "NO")
			if err2 != nil {
				return
			}
			return
		}

	}

	_, err3 := fmt.Fprintln(writer, strings.Join(result, "\n"))
	if err3 != nil {
		return
	}

}
