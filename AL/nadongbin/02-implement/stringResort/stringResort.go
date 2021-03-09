package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func stringResort(input string) string {
	charArr := strings.Split(input, "")

	str := []string{}
	sum := 0

	for _, char := range charArr {
		if char >= "A" && char <= "Z" {
			str = append(str, char)
		} else {
			num, _ := strconv.Atoi(char)
			sum += num
		}
	}

	sort.Strings(str)

	strSorted := strings.Join(str, "")

	return strSorted + strconv.Itoa(sum)
}

func main() {
	fmt.Println(stringResort("K1KA5CB7"))
	fmt.Println(stringResort("Z12DDV23"))
}
