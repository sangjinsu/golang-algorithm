package 위장

func solution(clothes [][]string) int {
	wear := make(map[string]int)

	for _, clothe := range clothes {
		if _, ok := wear[clothe[1]]; !ok {
			wear[clothe[1]] = 1
		} else {
			wear[clothe[1]]++
		}
	}

	answer := 1
	for _, v := range wear {
		answer *= v + 1
	}
	answer--
	return answer
}
