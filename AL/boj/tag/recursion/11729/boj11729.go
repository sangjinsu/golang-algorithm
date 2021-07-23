package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(int(math.Pow(2, float64(n))-1)) + "\n")

	hanoi(n, 1, 3, 2, &sb)
	fmt.Fprint(writer, sb.String())
}

func hanoi(n, start, to, step int, sb *strings.Builder) {
	move := func(start, to int, sb *strings.Builder) {
		sb.WriteString(strconv.Itoa(start) + " " + strconv.Itoa(to) + "\n")
	}

	if n == 1 {
		move(start, to, sb)
	} else {
		hanoi(n-1, start, step, to, sb)
		move(start, to, sb)
		hanoi(n-1, step, to, start, sb)
	}
}
