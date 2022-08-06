package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	var result bytes.Buffer
	for i := 0; i < N; i++ {
		stars := strings.Repeat("*", i+1)
		result.WriteString(stars)
		result.WriteString("\n")
	}

	fmt.Fprint(writer, result.String())
}
