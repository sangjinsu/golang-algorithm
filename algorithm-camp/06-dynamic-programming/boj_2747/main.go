package main

import (
	"bufio"
	"fmt"
	"os"
)

var cache = make(map[int]int)

func fibo(n int) int {

	if v, ok := cache[n]; ok {
		return v
	}

	if n <= 1 {
		cache[n] = n
		return n
	}

	result := fibo(n-1) + fibo(n-2)
	cache[n] = result
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	fmt.Fprintln(writer, fibo(n))
}
