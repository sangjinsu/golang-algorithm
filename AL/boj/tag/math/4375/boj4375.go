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

	for {
		var n int
		_, err := fmt.Fscan(reader, &n)
		if err != nil {
			break
		}

		for i, t := 1, 1; ; i++ {
			t %= n
			if t == 0 {
				fmt.Fprintf(writer, "%d\n", i)
				break
			}
			t = t*10 + 1
		}
	}
}
