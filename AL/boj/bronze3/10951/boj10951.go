package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			return
		}
	}(writer)

	for true {
		var a, b int
		_, err := fmt.Fscan(reader, &a, &b)
		if err == io.EOF {
			break
		}
		_, err = fmt.Fprintln(writer, a + b)
		if err != nil {
			return
		}
	}

}
