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
	defer writer.Flush()

	var chars string
	var m int
	fmt.Fscanln(reader, &chars)
	fmt.Fscanln(reader, &m)

	left := list.New()
	right := list.New()

	for _, char := range chars {
		left.PushBack(char)
	}

	for i := 0; i < m; i++ {
		instruction, _ := reader.ReadString('\n')
		strings.TrimSuffix(instruction, "\n")
		keyword := instruction[0]
		if 'L' == keyword && left.Len() > 0 {
			item := left.Back()
			char := item.Value
			left.Remove(item)
			right.PushFront(char)
		}

		if 'D' == keyword && right.Len() > 0{
			item := right.Front()
			char := item.Value
			right.Remove(item)
			left.PushBack(char)
		}

		if 'B' == keyword && left.Len() > 0 {
			item := left.Back()
			left.Remove(item)
		}

		if 'P' == keyword {
			left.PushBack(instruction[2])
		}
	}

	for i := left.Front(); i != nil ; i = i.Next() {
		fmt.Fprint(writer, fmt.Sprintf("%c", i.Value))
	}

	for i := right.Front(); i != nil ; i = i.Next() {
		fmt.Fprint(writer, fmt.Sprintf("%c", i.Value))
	}

}
