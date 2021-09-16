package main

import (
	"bufio"
	"fmt"
	"os"
)

type year struct {
	e int
	s int
	m int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n1, n2, n3 int
	fmt.Fscanf(reader, "%d %d %d", &n1, &n2, &n3)
	there := year{n1, n2, n3}
	here := year{1, 1, 1}

	y := 1
	for {
		if there.e == here.e && there.s == here.s && there.m == here.m {
			break
		}
		here.e = (here.e + 1) % 16
		if here.e == 0 {
			here.e = 1
		}
		here.s = (here.s + 1) % 29
		if here.s == 0 {
			here.s = 1
		}
		here.m = (here.m + 1) % 20
		if here.m == 0 {
			here.m = 1
		}
		y++
	}

	fmt.Fprint(writer, y)
}
