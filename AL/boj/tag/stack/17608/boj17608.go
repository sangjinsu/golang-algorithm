package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	var sticks []int

	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		var stick int
		fmt.Fscan(reader, &stick)
		sticks = append(sticks, stick)
	}

	highest := 0
	count := 0
	for len(sticks) > 0 {
		stick := sticks[len(sticks)-1]
		sticks = sticks[:len(sticks)-1]
		if highest < stick {
			count++
			highest = stick
		}
	}

	fmt.Fprint(writer, count)
}