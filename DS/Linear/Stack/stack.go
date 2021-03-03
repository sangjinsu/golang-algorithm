package main

import "fmt"

type node struct {
	data int
	next *node
}

type stack struct {
	top *node
}

func (s *stack) pop() (int, error) {
	if s.top == nil {
		return -1, fmt.Errorf("top is empty")
	}

	item := s.top.data
	s.top = s.top.next

	return item, nil
}

func (s *stack) push(item int) {
	node := node{data: item}
	node.next = s.top
	s.top = &node
}

func (s *stack) peek() (int, error) {
	if s.top == nil {
		return -1, fmt.Errorf("top is empty")
	}
	return s.top.data, nil
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

func main() {
	s := stack{}
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.peek())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
