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

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	bookSales := make(map[string]int, n)
	bookTitles := make([]string, n)
	for i := 0; i < n; i++ {
		var bookTitle string
		fmt.Fscanf(reader, "%s\n", &bookTitle)
		bookSales[bookTitle]++
		bookTitles = append(bookTitles, bookTitle)
	}

	sort.Strings(bookTitles)
	mostSale, bestSeller := 0, ""
	for _, title := range bookTitles {
		if bookSales[title] > mostSale {
			mostSale = bookSales[title]
			bestSeller = title
		}
	}

	fmt.Fprintf(writer, "%s", bestSeller)
}
