package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Meeting struct {
	start int
	end   int
}

type Meetings []Meeting

func (m Meetings) Len() int {
	return len(m)
}

func (m Meetings) Less(i, j int) bool {
	if m[i].end < m[j].end {
		return true
	} else if m[i].end == m[j].end {
		return m[i].start < m[j].start
	}
	return false
}

func (m Meetings) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)

	meetings := make(Meetings, N)
	for i := 0; i < N; i++ {
		var start, end int
		fmt.Fscanf(reader, "%d %d\n", &start, &end)
		meetings[i] = Meeting{start: start, end: end}
	}

	sort.Sort(meetings)

	cnt := 1
	end := meetings[0].end
	for i := 1; i < N; i++ {
		if end <= meetings[i].start {
			cnt++
			end = meetings[i].end
		}
	}
	fmt.Fprint(writer, cnt)
}
