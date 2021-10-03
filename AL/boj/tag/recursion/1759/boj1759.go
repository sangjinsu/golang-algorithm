package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	C       int
	L       int
	letters []string
	reader  *bufio.Reader
	writer  *bufio.Writer
)

func recursion(idx int, word string) {
	if len(word) == L {
		var cnt int
		for _, letter := range word {
			for _, r := range []rune{'a', 'e', 'i', 'o', 'u'} {
				if letter == r {
					cnt++
				}
			}
		}
		if cnt >= 1 && L-cnt >= 2 {
			fmt.Fprintln(writer, word)
		}
		return
	}
	if idx >= C {
		return
	}
	recursion(idx+1, word+letters[idx])
	recursion(idx+1, word)
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &L, &C)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	letters = strings.Split(line, " ")

	sort.Strings(letters)
	recursion(0, "")
}
