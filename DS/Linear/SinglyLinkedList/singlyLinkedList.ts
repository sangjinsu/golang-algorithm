class node {
  public next: node | null = null
  constructor(public value: number) {}
}

class SinglyLinkedList {
  public length: number = 0
  public head: node | null = null
  public tail: node | null = null

  // 노드 추가
  push(value: number) {
    const newNode = new node(value)
    if (this.head === null) {
      this.head = newNode
      this.tail = newNode
    } else if (this.tail !== null) {
      this.tail.next = newNode
      this.tail = newNode
    }
    this.length++
    return this
  }

  // 마지막 노드 제거
  pop() {
    if (this.head === null) return undefined
    let current = this.head
    let newTail = current
    while (current.next) {
      newTail = current
      current = current.next
    }
    this.tail = newTail
    this.tail.next = null
    this.length--
    if (this.length === 0) {
      this.head = null
      this.tail = null
    }
    return current
  }

  // 맨 앞 노드 제거
  shift() {
    if (this.head === null) return undefined
    const currentHead = this.head
    this.head = currentHead.next
    this.length--
    if (this.length === 0) {
      this.tail = null
    }
    return currentHead
  }

  // 맨 앞 노드 추가
  unshift(value: number) {
    const newNode = new node(value)
    if (this.head === null) {
      this.head = newNode
      this.tail = newNode
    } else {
      newNode.next = this.head
      this.head = newNode
    }
    this.length++
    return this
  }

  // 노드 가져오기
  get(index: number) {
    if (index < 0 || index >= this.length) return null

    let counter = 0
    let current = this.head as node
    while (counter !== index) {
      current = current.next as node
      counter++
    }

    return current
  }

  // 노드 값 변경
  set(index: number, value: number) {
    let foundNode = this.get(index)
    if (foundNode !== null) {
      foundNode.value = value
      return true
    }
    return false
  }

  // 노드 삽입
  insert(index: number, value: number) {
    if (index < 0 || index > this.length) return false
    if (index === this.length) return !!this.push(value)
    if (index === 0) return !!this.unshift(value)

    const newNode = new node(value)

    let prev = this.get(index - 1) as node
    let nextNode = prev.next as node
    prev.next = newNode
    newNode.next = nextNode
    this.length++
    return true
  }

  // 노드 제거
  remove(index: number) {
    if (index < 0 || index > this.length) return undefined
    if (index === this.length) return this.pop()
    if (index === 0) return this.shift()

    let prev = this.get(index - 1) as node
    const removed = prev.next as node
    prev.next = removed.next
    this.length--
    return removed
  }

  reverse() {
    if (this.head === null) {
      return this
    }
    let node = this.head as node
    this.head = this.tail
    this.tail = node
    let next: any
    let prev = null
    for (let i = 0; i < this.length; i++) {
      next = node.next
      node.next = prev
      prev = node
      node = next
    }
    return this
  }
}
