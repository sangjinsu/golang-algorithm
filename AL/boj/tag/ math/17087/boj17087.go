package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, s int
	fmt.Fscanln(reader, &n, &s)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	var positions []int
	inputs := strings.Split(input, " ")
	for _, input := range inputs {
		pos, _ := strconv.Atoi(input)
		positions = append(positions, pos)
	}

	var distances []int
	for _, position := range positions {
		distances = append(distances, int(math.Abs(float64(position-s))))
	}

	if n == 1 {
		fmt.Fprintln(writer, distances[0])
	} else {
		result := gcd(distances[0], distances[1])
		for i := 2; i < len(distances); i++ {
			result = gcd(result, distances[i])
		}
		fmt.Fprintln(writer, result)
	}
}
