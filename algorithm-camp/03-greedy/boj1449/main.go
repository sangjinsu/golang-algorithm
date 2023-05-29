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

	var N, L int
	fmt.Fscanf(reader, "%d %d\n", &N, &L)

	pipe := make([]bool, 1001)

	var hole int
	for i := 0; i < N; i++ {
		fmt.Fscanf(reader, "%d", &hole)
		pipe[hole] = true
	}

	var result int
	for i, hole := range pipe {
		if hole {
			result++
			for j := i; j < i+L; j++ {
				if j > 1000 {
					break
				}
				pipe[j] = false
			}
		}
	}

	fmt.Fprintf(writer, "%d", result)
}
