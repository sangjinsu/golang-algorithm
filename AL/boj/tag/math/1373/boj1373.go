package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func binaryToOctal(binary string) string {
	octal := 0
	num := 1
	for i := len(binary) - 1; i >= 0; i-- {
		b, _ := strconv.Atoi(string(binary[i]))
		octal += b * num
		num *= 2
	}

	return strconv.Itoa(octal)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var binary string
	fmt.Fscanln(reader, &binary)

	var reversedOctalBuffer bytes.Buffer
	for i := len(binary); i >= 0; i -= 3 {

		if i-3 < 0 {
			if binary[:i] != "" {
				reversedOctalBuffer.WriteString(binaryToOctal(binary[:i]))
			}
		} else {
			reversedOctalBuffer.WriteString(binaryToOctal(binary[i-3 : i]))
		}
	}

	var octal bytes.Buffer
	reversedOctal := reversedOctalBuffer.String()
	for i := len(reversedOctal) - 1; i >= 0; i-- {
		octal.WriteString(string(reversedOctal[i]))
	}
	fmt.Fprint(writer, octal.String())

}
