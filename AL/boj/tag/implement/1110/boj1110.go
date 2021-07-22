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
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	k := n
	count := 0
	for true {
		count++
		nums := strings.Split(strconv.Itoa(k), "")
		sum := 0
		for _, num := range nums {
			number, _ := strconv.Atoi(num)
			sum += number
		}
		sumArr := strings.Split(strconv.Itoa(sum), "")
		k, _ = strconv.Atoi(nums[len(nums) - 1] + sumArr[len(sumArr) - 1])

		if k == n {
			break
		}
	}

	fmt.Fprint(writer, count)
}
