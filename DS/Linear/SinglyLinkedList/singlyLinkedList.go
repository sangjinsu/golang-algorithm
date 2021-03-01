package main

import "fmt"

type Node struct {
	next  *Node
	value int
}

type SinglyLinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func (sl *SinglyLinkedList) push(value int) *SinglyLinkedList {
	node := &Node{value: value}
	if sl.head == nil {
		sl.head = node
		sl.tail = node
	} else {
		sl.tail.next = node
		sl.tail = node
	}
	sl.length++
	return sl
}

func (sl *SinglyLinkedList) pop() (*Node, error) {
	if sl.head == nil {
		return nil, fmt.Errorf("Head is not existing")
	}
	current := sl.head
	newTail := current

	for current.next != nil {
		newTail = current
		current = current.next
	}

	sl.tail = newTail
	sl.tail.next = nil
	sl.length--

	if sl.length == 0 {
		sl.head = nil
		sl.tail = nil
	}
	return current, nil
}

func (sl *SinglyLinkedList) get(index int) (*Node, error) {
	if index < 0 || index >= sl.length {
		return nil, fmt.Errorf("Reset Index")
	}
	counter := 0
	current := sl.head
	for counter != index {
		current = current.next
		counter++
	}
	return current, nil
}

func main() {
	sl := SinglyLinkedList{}
	sl.push(5)
	sl.push(10)
	node, err := sl.get(6)
	if err != nil {
		fmt.Println(err)
		fmt.Println(node)
	}
	fmt.Println(sl)
}
