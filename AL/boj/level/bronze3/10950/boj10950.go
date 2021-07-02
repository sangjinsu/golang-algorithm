package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	var t, a, b int

	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	sc.Scan()
	t, _ = strconv.Atoi(sc.Text())

	for t > 0 {
		sc.Scan()
		input := strings.Split(sc.Text(), " ")
		a, _ = strconv.Atoi(input[0])
		b, _ = strconv.Atoi(input[1])
		_, err := writer.WriteString(strconv.Itoa(a + b) + "\n")
		if err != nil {
			return
		}
		func(writer *bufio.Writer) {
			err := writer.Flush()
			if err != nil {
				return
			}
		}(writer)
		t--
	}

}
