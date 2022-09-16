package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var white = 0
var blue = 0

func recursion(table [][]int, ny, nx, size int) {

	if checkColor(table, ny, nx, size) {
		color := table[ny][nx]
		if color == 1 {
			blue++
		} else {
			white++
		}
		return
	}

	halfSize := size / 2
	recursion(table, ny, nx, halfSize)
	recursion(table, ny+halfSize, nx, halfSize)
	recursion(table, ny, nx+halfSize, halfSize)
	recursion(table, ny+halfSize, nx+halfSize, halfSize)
}

func checkColor(table [][]int, ny int, nx int, size int) bool {
	color := table[ny][nx]
	for ty := ny; ty < ny+size; ty++ {
		for tx := nx; tx < nx+size; tx++ {
			if color != table[ty][tx] {
				return false
			}
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	table := make([][]int, N)
	for i := 0; i < N; i++ {
		line, _ := reader.ReadString('\n')
		values := strings.Fields(line)
		nums := make([]int, N)
		for j, value := range values {
			num, _ := strconv.Atoi(value)
			nums[j] = num
		}
		table[i] = nums
	}

	recursion(table, 0, 0, N)
	fmt.Fprintln(writer, white)
	fmt.Fprint(writer, blue)
}
