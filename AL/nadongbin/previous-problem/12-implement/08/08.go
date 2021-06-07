package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputs := []string{
		"K1KA5CB7", "AJKDLSI412K4JSJ9D",
	}

	for i := 0; i < len(inputs); i++ {
		solution(inputs[i])
	}

}

func solution(input string) {
	var result string
	var alpabet []string
	var sum int

	for _, v := range input {
		if "A" <= string(v) && string(v) <= "Z" {
			alpabet = append(alpabet, string(v))
		} else {
			n, _ := strconv.Atoi(string(v))
			sum += n
		}
	}

	sort.Strings(alpabet)
	result = strings.Join(alpabet, "")
	result += strconv.Itoa(sum)

	fmt.Println(result)
}
