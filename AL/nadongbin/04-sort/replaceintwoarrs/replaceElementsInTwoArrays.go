package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	fmt.Scanln(&n, &k)

	arr1 := make([]int, n)
	arr2 := make([]int, n)

	for i := 0; i < 2; i++ {
		var str string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		str = scanner.Text()

		strArr := strings.Split(str, " ")
		for j := 0; j < len(strArr); j++ {
			if i == 0 {
				arr1[j], _ = strconv.Atoi(strArr[j])
			} else {
				arr2[j], _ = strconv.Atoi(strArr[j])
			}
		}
	}

	sort.Ints(arr1)
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] > arr2[j]
	})

	for i := 0; i < k; i++ {
		if arr1[i] < arr2[i] {
			arr1[i], arr2[i] = arr2[i], arr1[i]
		} else {
			break
		}
	}

	sum := 0
	for _, value := range arr1 {
		sum += value
	}

	fmt.Println(sum)

}
