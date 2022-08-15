package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var chars string
	fmt.Fscanln(reader, &chars)

	var result []rune
	for _, c := range chars {
		if unicode.IsUpper(c) {
			result = append(result, unicode.ToLower(c))
		} else {
			result = append(result, unicode.ToUpper(c))
		}
	}

	fmt.Fprint(writer, string(result))
}
