package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	maxNum    uint64 = 0
	minNum    uint64 = 9999999999
	maxResult string
	minResult string
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	reader.Scan()
	K, _ := strconv.Atoi(reader.Text())

	reader.Scan()
	signs := strings.Fields(reader.Text())

	visited := make([]bool, 10)
	for num := 0; num < 10; num++ {
		visited[num] = true
		recursion(K, 0, strconv.Itoa(num), visited, signs)
		visited[num] = false
	}

	fmt.Fprintf(writer, "%v\n%v\n", maxResult, minResult)
}

func recursion(k int, idx int, value string, visited []bool, signs []string) {
	if idx == k {

		num, _ := strconv.ParseUint(value, 10, 64)
		if maxNum < num {
			maxResult = value
			maxNum = num
		}
		if minNum > num {
			minResult = value
			minNum = num
		}
		return
	}

	beforeIndex := 0

	if len(value) >= 2 {
		beforeIndex = len(value) - 1
	}

	if signs[idx] == "<" {
		for num := 0; num < 10; num++ {
			if visited[num] == false {
				v, _ := strconv.Atoi(string(value[beforeIndex]))
				if v < num {
					visited[num] = true
					recursion(k, idx+1, value+strconv.Itoa(num), visited, signs)
					visited[num] = false
				}
			}
		}
	} else {
		for num := 0; num < 10; num++ {
			if visited[num] == false {
				v, _ := strconv.Atoi(string(value[beforeIndex]))
				if v > num {
					visited[num] = true
					recursion(k, idx+1, value+strconv.Itoa(num), visited, signs)
					visited[num] = false
				}
			}
		}
	}
}
