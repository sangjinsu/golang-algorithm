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

	var score string
	fmt.Fscanln(reader, &score)

	scores := map[string]float64{
		"A0": 4.0, "B0": 3.0, "C0": 2.0, "D0": 1.0, "F": 0.0,
		"A+": 4.3, "B+": 3.3, "C+": 2.3, "D+": 1.3,
		"A-": 3.7, "B-": 2.7, "C-": 1.7, "D-": 0.7,
	}

	if result, ok := scores[score]; ok {
		fmt.Fprintf(writer, "%.1f\n", result)
	}
}
