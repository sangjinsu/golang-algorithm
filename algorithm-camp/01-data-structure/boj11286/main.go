package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority == pq[j].priority {
		return pq[i].value < pq[j].value
	}

	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscanf(reader, "%d\n", &num)

		if num != 0 {
			item := &Item{
				value:    num,
				priority: int(math.Abs(float64(num))),
				index:    i,
			}
			heap.Push(&pq, item)
		} else {
			if pq.Len() > 0 {
				fmt.Fprintf(writer, "%d\n", heap.Pop(&pq).(*Item).value)
			} else {
				fmt.Fprintf(writer, "%d\n", 0)
			}
		}
	}
}
