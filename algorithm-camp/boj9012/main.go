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

	var T int
	fmt.Fscanf(reader, "%d\n", &T)

	for i := 0; i < T; i++ {
		var cnt int
		var brackets string
		fmt.Fscanf(reader, "%s\n", &brackets)

		result := "YES"
		for _, b := range brackets {
			if string(b) == "(" {
				cnt++
			} else {
				if cnt == 0 {
					result = "NO"
					break
				}
				cnt--
			}
		}
		if cnt != 0 {
			result = "NO"
		}

		fmt.Fprintln(writer, result)
	}
}
