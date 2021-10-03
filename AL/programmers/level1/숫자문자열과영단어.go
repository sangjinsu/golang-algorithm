package level1

import "strconv"

func solution(s string) int {
	table := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	var nums, temp string
	for _, letter := range s {
		if _, err := strconv.Atoi(string(letter)); err == nil {
			nums += string(letter)
		} else {
			temp += string(letter)
			for key, value := range table {
				if temp == key {
					nums += value
					temp = ""
				}
			}
		}
	}
	result, _ := strconv.Atoi(nums)
	return result
}
