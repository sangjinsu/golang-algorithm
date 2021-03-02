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
		return nil, fmt.Errorf("check index please")
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

// 중복값 삭제
func (sl *SinglyLinkedList) removeDups() {
	node := sl.head
	for node != nil && node.next != nil {
		r := node
		for r.next != nil {
			if node.value == r.next.value {
				r.next = r.next.next
			} else {
				r = r.next
			}
		}
		node = node.next
	}
}

func sumLists(l1 SinglyLinkedList, l2 SinglyLinkedList) {

	result := SinglyLinkedList{}

	node1, err := l1.get(0)
	if err != nil {
		return
	}

	node2, err := l2.get(0)
	if err != nil {
		return
	}

	length := 0
	sum := 0
	carry := 0
	for node1 != nil && node2 != nil {
		sum = 0
		sum += carry
		sum += node1.value
		sum += node2.value

		result.push(sum % 10)

		if sum > 10 {
			carry = 1
		} else {
			carry = 0
		}

		length++

		node1 = node1.next
		node2 = node2.next
	}

	if length < l1.length {
		for node1 != nil {
			result.push(node1.value + carry)
			carry = 0
			node1 = node1.next
		}
	} else if length < l2.length {
		for node2 != nil {
			result.push(node2.value + carry)
			carry = 0
			node2 = node2.next
		}
	} else if carry != 0 {
		result.push(carry)
	}

	for i := 0; i < result.length; i++ {
		node, err := result.get(i)
		if err != nil {
			return
		}
		fmt.Printf("%v ", node.value)
	}
}

func getIntersection(l1, l2 SinglyLinkedList) (*Node, error) {
	var n1 *Node
	var n2 *Node
	if l1.length > l2.length {
		n1, _ = l1.get(l1.length - l2.length)
		n2, _ = l2.get(0)

	} else if l1.length < l2.length {
		n2, _ = l2.get(l2.length - l1.length)
		n1, _ = l1.get(0)
	}
	for n1 != nil && n2 != nil {
		if n1.value == n2.value {
			return n1, nil
		}
		n1 = n1.next
		n2 = n2.next
	}
	return nil, fmt.Errorf("Not intersection")
}

func findLoop(head *Node) {
	fast := head
	slow := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.next == nil {
		return
	}

	slow = head
	for fast != slow {
		slow = slow.next
		fast = fast.next
	}

	fmt.Println(fast)
}

func main() {
	// Linked List Digit 합산 알고리즘 in Go
	l1 := SinglyLinkedList{}
	l1.push(5)
	l1.push(5)
	l1.push(5)
	l1.push(8)

	l2 := SinglyLinkedList{}
	l2.push(3)
	l2.push(9)
	l2.push(9)

	sumLists(l1, l2)
	fmt.Println()

	// 교차점 찾기
	l3 := SinglyLinkedList{}
	l3.push(1)
	l3.push(2)
	l3.push(3)
	l3.push(7)
	l3.push(8)

	l4 := SinglyLinkedList{}
	l4.push(4)
	l4.push(5)
	l4.push(7)
	l4.push(8)

	node, _ := getIntersection(l3, l4)
	fmt.Println(node.value)

	// 루프 찾기

	l5 := SinglyLinkedList{}
	l5.push(1) // node0
	l5.push(2) // node1
	l5.push(3)
	l5.push(4)
	l5.push(5)
	l5.push(6)
	l5.push(7)
	l5.push(8) // node7
	node7, _ := l5.get(7)
	node4, _ := l5.get(2)
	node7.next = node4

	findLoop(l5.head)
}
