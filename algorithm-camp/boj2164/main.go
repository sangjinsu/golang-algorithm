package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N uint32
	fmt.Fscanf(reader, "%d\n", &N)

	cards := list.New().Init()
	for i := N; i > 0; i-- {
		cards.PushBack(i)
	}

	for cards.Len() > 1 {
		card := cards.Back()
		cards.Remove(card)
		card = cards.Back()
		cards.MoveToFront(card)
	}

	card := cards.Back()
	fmt.Fprint(writer, card.Value)
}
