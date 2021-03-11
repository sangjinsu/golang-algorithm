package main

import (
	"container/list"
	"fmt"
)

func main() {
	stack := list.New()
	stack.PushBack(5)
	stack.PushBack(2)
	stack.PushBack(3)
	stack.PushBack(7)

	e := stack.Back()
	stack.Remove(e)

	stack.PushBack(1)
	stack.PushBack(2)

	e = stack.Back()
	stack.Remove(e)

	for stack.Len() > 0 {
		e := stack.Back()
		fmt.Println(e.Value)
		stack.Remove(e)
	}
}
