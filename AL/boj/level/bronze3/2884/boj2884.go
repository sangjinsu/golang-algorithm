package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	sb := strings.Builder{}

	var hour, minute int
	_, err := fmt.Fscanln(reader, &hour, &minute)
	if err != nil {
		return
	}

	minute -= 45

	if minute < 0 {
		minute = 60 + minute
		hour--
	}

	if hour < 0 {
		hour = 23
	}
	
	sb.WriteString(strconv.Itoa(hour) + " " + strconv.Itoa(minute))

	_, err = fmt.Fprint(writer, sb.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Frint: %v\n", err)
	}
}
