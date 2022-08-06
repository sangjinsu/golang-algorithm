package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	var result bytes.Buffer
	for i := 1; i < 10; i++ {
		str := fmt.Sprintf("%d * %d = %d\n", N, i, i*N)
		result.WriteString(str)
	}
	fmt.Fprint(writer, result.String())
}
