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

	notHearSee := make(map[string]bool)

	for i := 0; i < n; i++ {
		var name string
		fmt.Fscanln(reader, &name)
		notHearSee[name] = true
	}

	var result []string
	for i := 0; i < m; i++ {
		var name string
		fmt.Fscanln(reader, &name)
		_, exist := notHearSee[name]
		if exist {
			result = append(result, name)
		}
	}

	sort.Strings(result)
	fmt.Fprintln(writer, len(result))
	for _, s := range result {
		fmt.Fprintln(writer, s)
	}
}
