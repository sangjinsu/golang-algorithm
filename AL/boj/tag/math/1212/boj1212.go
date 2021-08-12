package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func octalToBinary(o string) string {
	if o == "0"{
		return o
	}
	octal, _ := strconv.ParseInt(o, 10, 64)
	var result string
	for octal != 0 {
		octal, result = octal/2, strconv.Itoa(int(octal%2))+result
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var octal string
	fmt.Fscanln(reader, &octal)

	result := strings.Builder{}
	for i, o := range octal {
		b := octalToBinary(string(o))
		if i > 0 {
			for len(b) < 3 {
				b = "0" + b
			}
		}
		result.WriteString(b)
	}

	fmt.Fprint(writer, result.String())
}
