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

func (sl *SinglyLinkedList) shift() (*Node, error) {
	if sl.head == nil {
		return nil, fmt.Errorf("Head is not existing")
	}
	currentHead := sl.head
	sl.head = currentHead.next
	sl.length--
	if sl.length == 0 {
		sl.tail = nil
	}
	return currentHead, nil
}

func (sl *SinglyLinkedList) unshift(value int) *SinglyLinkedList {
	newNode := &Node{value: value}
	if sl.head == nil {
		sl.head = newNode
		sl.tail = newNode
	} else {
		newNode.next = sl.head
		sl.head = newNode
	}
	sl.length++
	return sl
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

func (sl *SinglyLinkedList) set(index, value int) bool {
	foundNode, err := sl.get(index)
	if err != nil {
		fmt.Println("Not found Node, please set other index")
		return false
	}
	foundNode.value = value
	return true
}

func (sl *SinglyLinkedList) insert(index, value int) bool {
	if index < 0 || index > sl.length {
		return false
	}
	if index == sl.length {
		sl.push(value)
		return true
	}
	if index == 0 {
		sl.unshift(value)
		return true
	}
	newNode := &Node{value: value}
	prev, err := sl.get(index - 1)
	if err != nil {
		return false
	}
	nextNode := prev.next
	newNode.next = nextNode
	sl.length++
	return true
}

func (sl *SinglyLinkedList) remove(index int) (*Node, error) {
	if index < 0 || index > sl.length {
		return sl.pop()
	}

	if index == sl.length {
		return sl.pop()
	}

	if index == 0 {
		return sl.shift()
	}

	prev, err := sl.get(index - 1)
	if err != nil {
		return nil, fmt.Errorf("index is not existing")
	}

	removed := prev.next
	prev.next = removed.next
	sl.length--
	return removed, nil
}
