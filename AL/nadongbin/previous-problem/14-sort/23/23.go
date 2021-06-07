package main

import (
	"fmt"
	"math"
	"sort"
)

type student struct {
	name string
	ko   int
	en   int
	ma   int
}

func main() {
	var n int
	fmt.Scanln(&n)

	students := make([]student, n)

	for i := 0; i < n; i++ {
		var name string
		var ko, en, ma int
		fmt.Scanln(&name, &ko, &en, &ma)
		students[i] = student{name: name, ko: ko, en: en, ma: ma}
	}

	sort.Slice(students, func(i, j int) bool {

		if students[i].ko == students[j].ko && students[i].en == students[j].en && students[i].ma == students[j].ma {
			length := int(math.Min(float64(len(students[i].name)), float64(len(students[j].name))))
			for k := 0; k < length; k++ {
				if students[i].name[k] < students[j].name[k] {
					return students[i].name[k] < students[j].name[k]
				}
			}
			return false
		}

		if students[i].ko == students[j].ko && students[i].en == students[j].en {
			return students[i].ma > students[j].ma
		}

		if students[i].ko == students[j].ko {
			return students[i].en < students[j].en
		}

		return students[i].ko > students[j].ko
	})

	for _, s := range students {
		fmt.Println(s.name)
	}
}
