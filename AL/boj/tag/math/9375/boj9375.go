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
	fmt.Fscanln(reader, &T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Fscanln(reader, &N)

		clothes := make(map[string]int)
		for i := 0; i < N; i++ {
			var wear, wearType string
			fmt.Fscanf(reader, "%s %s\n", &wear, &wearType)

			if _, ok := clothes[wearType]; ok {
				clothes[wearType]++
			} else {
				clothes[wearType] = 1
			}
		}

		result := 1
		for _, num := range clothes {
			result *= num + 1
		}
		result--

		fmt.Fprintln(writer, result)
	}
}
