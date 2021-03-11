package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()
	queue.PushBack(5)
	queue.PushBack(2)
	queue.PushBack(3)
	queue.PushBack(7)

	e := queue.Front()
	queue.Remove(e)

	queue.PushBack(1)
	queue.PushBack(4)

	e = queue.Front()
	queue.Remove(e)

	for queue.Len() > 0 {
		e := queue.Front()
		fmt.Println(e.Value)
		queue.Remove(e)
	}

}
