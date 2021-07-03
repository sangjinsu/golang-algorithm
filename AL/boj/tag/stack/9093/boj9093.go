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

	var t int
	var sentence string

	_, err := fmt.Fscanln(reader, &t)
	if err != nil {
		return
	}

	sb := strings.Builder{}
	stack := list.New()

	for i := 0; i < t; i++ {
		sentence, _ = reader.ReadString('\n')
		strings.TrimSuffix(sentence, "\n")

		for _, letter := range sentence {
			if string(letter) == "\n" || string(letter) == " " {
				for stack.Len() > 0 {
					item := stack.Back()
					sb.WriteString(string(item.Value.(int32)))
					stack.Remove(item)
				}
				sb.WriteRune(' ')
			} else {
				stack.PushBack(letter)
			}
		}
	}

	_, err3 := fmt.Fprint(writer, sb.String())
	if err3 != nil {
		return
	}
}
