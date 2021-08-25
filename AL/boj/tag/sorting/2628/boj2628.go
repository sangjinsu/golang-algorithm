package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	var k int
	fmt.Fscanln(reader, &k)

	dottedLines := make([][2]int, k)
	for i := 0; i < k; i++ {
		var axis, num int
		fmt.Fscanf(reader, "%d %d\n", &axis, &num)
		dottedLines[i] = [2]int{axis, num}
	}

	rowCut := []int{0}
	colCut := []int{0}
	for _, line := range dottedLines {
		if line[0] == 0 {
			rowCut = append(rowCut, line[1])
		} else {
			colCut = append(colCut, line[1])
		}
	}

	rowCut = append(rowCut, m)
	colCut = append(colCut, n)

	sort.Ints(rowCut)
	sort.Ints(colCut)

	var cutRows, cutCols []int
	for i := 1; i < len(rowCut); i++ {
		cutRows = append(cutRows, rowCut[i]-rowCut[i-1])
	}
	for i := 1; i < len(colCut); i++ {
		cutCols = append(cutCols, colCut[i]-colCut[i-1])
	}
	sort.Ints(cutRows)
	sort.Ints(cutCols)

	result := cutRows[len(cutRows)-1] * cutCols[len(cutCols)-1]
	fmt.Fprint(writer, result)
}
